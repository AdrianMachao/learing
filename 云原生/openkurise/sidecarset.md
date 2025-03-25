# sidecarset
针对sidecar管理
```
apiVersion: apps.kruise.io/v1alpha1
kind: SidecarSet
metadata:
  name: sidecarset
spec:
  selector:
    matchLabels:
      app: sample
  containers:
  - name: nginx
    image: nginx:alpine
  initContainers:
  - name: init-container
    image: busybox:latest
    command: [ "/bin/sh", "-c", "sleep 5 && echo 'init container success'" ]
  updateStrategy:
    type: RollingUpdate
  namespace: ns-1
```
- spec.selector 通过label的方式选择需要注入、更新的pod，支持matchLabels、MatchExpressions两种方式，详情请参考：https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/
- spec.containers 定义需要注入、更新的pod.spec.containers容器，支持完整的k8s container字段，详情请参考：https://kubernetes.io/docs/concepts/containers/
- spec.initContainers 定义需要注入的pod.spec.initContainers容器，支持完整的k8s initContainer字段，详情请参考：https://kubernetes.io/docs/concepts/workloads/pods/init-containers/
  - 注入initContainers容器默认基于container name升级排序
  - initContainers只支持注入，不支持pod原地升级
- spec.updateStrategy sidecarSet更新策略，type表明升级方式：
  - NotUpdate 不更新，此模式下只会包含注入能力
  - RollingUpdate 注入+滚动更新，包含了丰富的滚动更新策略，后面会详细介绍
spec.namespace sidecarset默认在k8s整个集群范围内生效，即对所有的命名空间生效（除了kube-system, kube-public），当设置该字段时，只对该namespace的pod生效


## 注入管理
sidecar 的注入只会发生在 Pod 创建阶段，并且只有 Pod spec 会被更新，不会影响 Pod 所属的 workload template 模板。 spec.containers除了默认的k8s container字段，还扩展了如下一些字段，来方便注入
```
apiVersion: apps.kruise.io/v1alpha1
kind: SidecarSet
metadata:
  name: sidecarset
spec:
  selector:
    matchLabels:
      app: sample
  containers:
    # default K8s Container fields
  - name: nginx
    image: nginx:alpine
    volumeMounts:
    - mountPath: /nginx/conf
      name: nginx.conf
    # extended sidecar container fields
    podInjectPolicy: BeforeAppContainer
    shareVolumePolicy:
      type: disabled | enabled
    transferEnv:
    - sourceContainerName: main
      envName: PROXY_IP
    - sourceContainerNameFrom:
        fieldRef:
          apiVersion: "v1"
          fieldPath: "metadata.labels['cName']"
        # fieldPath: "metadata.annotations['cName']"
      envName: TC
  volumes:
  - name: nginx.conf
    hostPath:
      path: /data/nginx/conf
```
- podInjectPolicy 定义container注入到pod.spec.containers中的位置
  - BeforeAppContainer(默认) 注入到pod原containers的前面
  - AfterAppContainer 注入到pod原containers的后面
- 数据卷共享
  - 共享指定卷：通过 spec.volumes 来定义 sidecar 自身需要的 volume，详情请参考：https://kubernetes.io/docs/concepts/storage/volumes/
  - 共享所有卷：通过 spec.containers[i].shareVolumePolicy.type = enabled | disabled 来控制是否挂载pod应用容器的卷，常用于日志收集等 sidecar，配置为 enabled 后会把应用容器中所有挂载点注入 sidecar 同一路经下(sidecar中本身就有声明的数据卷和挂载点除外）
- 环境变量共享
  - 可以通过 spec.containers[i].transferEnv 来从别的容器获取环境变量，会把名为 sourceContainerName 容器中名为 envName 的环境变量拷贝到本容器
  - sourceContainerNameFrom 支持 downwardAPI 来获取容器name，比如：metadata.name, metadata.labels['<KEY>'], metadata.annotations['<KEY>']

## 更新策略

SidecarSet不仅支持sidecar容器的原地升级，而且提供了非常丰富的升级、灰度策略。

### 分批发布
Partition 的语义是 保留旧版本 Pod 的数量或百分比，默认为 0。这里的 partition 不表示任何 order 序号。

如果在发布过程中设置了 partition:
- 如果是数字，控制器会将 (replicas - partition) 数量的 Pod 更新到最新版本。
- 如果是百分比，控制器会将 (replicas * (100% - partition)) 数量的 Pod 更新到最新版本。

```
apiVersion: apps.kruise.io/v1alpha1
kind: SidecarSet
metadata:
  name: sidecarset
spec:
  # ...
  updateStrategy:
    type: RollingUpdate
    partition: 90
```
假设该SidecarSet关联的pod数量是100个，则本次升级只会升级10个，保留90个。

### 最大不可用数量

### 金丝雀发布
对于有金丝雀发布需求的业务，可以通过strategy.selector来实现。方式：对于需要率先金丝雀灰度的pod打上固定的labels[canary.release] = true，再通过strategy.selector.matchLabels来选中该pod
```
apiVersion: apps.kruise.io/v1alpha1
kind: SidecarSet
metadata:
  name: sidecarset
spec:
  # ...
  updateStrategy:
    type: RollingUpdate
    selector:
      matchLabels:
        canary.release: "true"
```


### 发布顺序控制

## 热升级
SidecarSet原地升级会先停止旧版本的容器，然后创建新版本的容器。这种方式更加适合不影响Pod服务可用性的sidecar容器，比如说：日志收集Agent。

但是对于很多代理或运行时的sidecar容器，例如Istio Envoy，这种升级方法就有问题了。Envoy作为Pod中的一个代理容器，代理了所有的流量，如果直接重启，Pod服务的可用性会受到影响。如果需要单独升级envoy sidecar，就需要复杂的grace终止和协调机制。所以我们为这种sidecar容器的升级提供了一种新的解决方案。
```
apiVersion: apps.kruise.io/v1alpha1
kind: SidecarSet
metadata:
  name: hotupgrade-sidecarset
spec:
  selector:
    matchLabels:
      app: hotupgrade
  containers:
  - name: sidecar
    image: openkruise/hotupgrade-sample:sidecarv1
    imagePullPolicy: Always
    lifecycle:
      postStart:
        exec:
          command:
          - /bin/sh
          - /migrate.sh
    upgradeStrategy:
      upgradeType: HotUpgrade
      hotUpgradeEmptyImage: openkruise/hotupgrade-sample:empty
```

- upgradeType: HotUpgrade代表该sidecar容器的类型是hot upgrade，将执行热升级方案
- hotUpgradeEmptyImage: 当热升级sidecar容器时，业务必须要提供一个empty容器用于热升级过程中的容器切换。empty容器同sidecar容器具有相同的配置（除了镜像地址），例如：command, lifecycle, probe等，但是它不做任何工作。
- lifecycle.postStart: 状态迁移，该过程完成热升级过程中的状态迁移，该脚本需要由业务根据自身的特点自行实现，例如：nginx热升级需要完成Listen FD共享以及流量排水（reload）

热升级特性总共包含以下两个过程：
1. Pod创建时，注入热升级容器
2. 原地升级时，完成热升级流程





