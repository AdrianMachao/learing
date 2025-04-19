# 概念
## gang
gang是一组调度的pod，需要满足最小pod数调度成功才可以成功

## ganggroup
ganggroup是一组gang，需要所有的gang满足要求才能运行？

## 运行模式
由于会存在死锁问题，当集群资源不足，gang1成功了2/5个，gang2成功了2/5个，需要三个以上才能调度m此时gang1与gang2成为死锁，都无法运行，但其实资源是能够运行一个gang的，此时会有strict模式和非srict模式
strict模式
如果有一个pod调度失败会立刻释放已经成功申请的pod
non-strict模式
pod调度失败会继续等待调度


# 实现
## preEnqueue
## queueSort
排序，按照gangGroup、gang、prority、createTime排序
## preFilter
预选，查看是否有满足资源需求的pod
resosurceStatisfield：是否已经满足调度
schedulerValid：周期是否有效
schedulerCycle：周期数
## postfilter

调度失败
strict模式：scheduleCycleValid设置无效，会释放gangGroup下已经调度成功的pod
非strict模式不做任何操作
## unreserve
超时或者bind失败
删除assume pod
strict模式：scheduleCycleValid设置无效，会释放gangGroup下已经调度成功的pod

## permit
准入阶段，pod调度成功但gang不满足要求时会在permit阶段等待，不会立刻进入bind阶段

## postbind
bind
resourceStatisfield
schedulerCycle
schedulerValid
