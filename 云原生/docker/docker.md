# 架构
![docker](./images/docker.jpg)

# 组件
1.docker
2.dockerd
3.containerd
4.containerd-shim
5.runC

## docker
docker 的命令行工具，是给用户和 docker daemon 建立通信的客户端。

## dockerd
dockerd 是 docker 架构中一个常驻在后台的系统进程，称为 docker daemon，dockerd 实际调用的还是 containerd 的 api 接口（rpc 方式实现）,docker daemon 的作用主要有以下两方面：

接收并处理 docker client 发送的请求
管理所有的 docker 容器

有了 containerd 之后，dockerd 可以独立升级，以此避免之前 dockerd 升级会导致所有容器不可用的问题。

## containerd
containerd 是 dockerd 和 runc 之间的一个中间交流组件，docker 对容器的管理和操作基本都是通过 containerd 完成的。containerd 的主要功能有：

- 容器生命周期管理
- 日志管理
- 镜像管理
- 存储管理
- 容器网络接口及网络管理

## containerd-shim
containerd-shim 是一个真实运行容器的载体，每启动一个容器都会起一个新的 containerd-shim 的一个进程， 它直接通过指定的三个参数：容器 id，boundle 目录（containerd 对应某个容器生成的目录，一般位于：/var/run/docker/libcontainerd/containerID，其中包括了容器配置和标准输入、标准输出、标准错误三个管道文件），运行时二进制（默认为 runC）来调用 runc 的 api 创建一个容器，上面的 docker 进程图中可以直观的显示。其主要作用是：

它允许容器运行时(即 runC)在启动容器之后退出，简单说就是不必为每个容器一直运行一个容器运行时(runC)
即使在 containerd 和 dockerd 都挂掉的情况下，容器的标准 IO 和其它的文件描述符也都是可用的
向 containerd 报告容器的退出状态

有了它就可以在不中断容器运行的情况下升级或重启 dockerd，对于生产环境来说意义重大。
运行是二进制（默认为 runc）来调用 runc 的 api 创建一个容器（比如创建容器：最后拼装的命令如下：runc create 。。。。。）

## runc
简单的说，runC 是一个命令行工具，用来运行按照 OCI 标准格式打包过的应用

启动一个容器的过程如下：

用户在命令行执行 docker run -itd busybox 由 docker client 通过 grpc 将指令传给 dockerd
docker daemon 请检查本机是否存在 docker 镜像文件，如果有继续往下执行
dockerd 会向 host os 请求创建容器
linux 会创建一个空的容器(cgroup namespace),并启动 containerd-shim 进程。
containerd-shim 拿到三个参数(容器 id，boundle 目录，运行时二进制文件 runc )来调用 runC 的 api
runC 提取镜像文件，生成容器配置文件，然后启动容器

# 内部通信图
![interactive](./images/interactive.jpg)

# cgroup
网络
挂载点
IPC
进程
用户
