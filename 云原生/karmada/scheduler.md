# 介绍
karmada是一个云原生多集群、多云编排工具，

## 调度
![karmada-scheduler](./images/karmada-scheduler.awebp)

karmada scheduler 在调度每个 k8s 原生 API 资源对象（包含 CRD 资源）时，会逐个调用各扩展点上的插件：
1.filter 扩展点上的调度算法插件将不满足 propagation policy 的成员集群过滤掉 karmada scheduler 对每个考察中的成员集群调用每个插件的 Filter 方法，该方法都能返回一个 Result 对象表示该插件的调度结果，其中的 code 代表待下发资源是否能调度到某个成员集群上，reason 用来解释这个结果，err 包含调度算法插件执行过程中遇到的错误.
最终按照第二步的评分高低选择成员集群作为调度结果。目前 karmada 的调度算法插件：
- APIInstalled: 用于检查资源的 API(CRD)是否安装在目标集群中。
- ClusterAffinity: 用于检查资源选择器是否与集群标签匹配。
- SpreadConstraint: 用于检查 Cluster.Spec 中的 spread 属性即 Provider/Zone/Region 字段。
- TaintToleration: 用于检查传播策略是否容忍集群的污点。
- ClusterLocality 是一个评分插件，为目标集群进行评分。
score 扩展点上的调度算法插件为每个经过上一步过滤的集群计算评分 karmada scheduler 对每个经过上一步过滤的成员集群调用每个插件的 Score 方法，该方法都能返回一个 int64 类型的评分结果。


这种着重说一下调度功能，调度主要涉及三个组件estimator、scheduler、descheduelr

### 触发时机
可以看到这里实现了三种场景的调度：

- 分发资源时选择目标集群的规则变了

- 副本数变了，即扩缩容调度

- 故障恢复调度，当被调度的成员集群状态不正常时会触发重新调度

### estimator
用于评估每个集群的资源，karmada scheduler-estimator 评估了以下资源：
- cpu
- memory
- ephemeral-storage
- 其他标量资源：（1）扩展资源，例如：requests.nvidia.com/gpu: 4（2）kubernetes.io/下原生资源（3）hugepages- 资源（4）attachable-volumes- 资源

### descheduler
重调度功能，定时检测集群调度后资源不足无法成功启动的情况，调度策略为动态划分（dynamic division）时才会生效，karmada-descheduler 将每隔一段时间检测一次所有部署，默认情况下每 2 分钟检测一次
每个周期中，它会通过调用 karmada-scheduler-estimator 找出部署在目标调度集群中有多少不可调度的副本，然后更新 ResourceBinding 资源的 Clusters[i].Replicas 字段，并根据当前情况触发 karmada-scheduler 执行“Scale Schedule”

### 调度策略
描述在多集群中有哪些调度策略
#### dulipcated
复制功能，每个集群部署的副本数目相同

#### divided
##### Wighted
静态划分策略，按照比例在集群中划分不同的策略，例如cluster1:cluster2=1:2

##### Aggregated
动态策略，也分两种，一种是精良部署在一个集群中，这里需要用到estimator来评估集群资源，但不能保证一定满足需求，所以会用到descheduler来触发重调度

#### 其他
例如按照可用区、机房等策略调度

## 容灾多活

