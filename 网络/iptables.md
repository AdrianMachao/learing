# iptables
iptables 是 Linux 内核中的防火墙软件 netfilter 的管理工具，位于用户空间，同时也是 netfilter 的一部分。Netfilter 位于内核空间，不仅有网络地址转换的功能，也具备数据包内容修改、以及数据包过滤等防火墙功能
![iptables](./images/iptables-flowchart.png)


## 表
raw 用于配置数据包，raw 中的数据包不会被系统跟踪。
filter 是用于存放所有与防火墙相关操作的默认表。
nat 用于 网络地址转换（例如：端口转发）。
mangle 用于对特定数据包的修改（参考损坏数据包）。
security 用于强制访问控制 网络规则

| Rules         | raw | filter | nat | mangle | security |
| ------------  | --- | -------| --- | ------ | -------- |
| PREROUTING    | ✓   |        |✓    |✓       |          |
| INPUT         |     |✓       |     |✓       |          |
| OUTPUT        | ✓   |✓       |✓    |✓       |          |
| POSTROUTING   |     |        |✓    |✓       |          |
| FORWARD       |     |✓       |     |✓       |          |

## iptables规则解释
查看 istio-proxy 容器中的默认的 iptables 规则，默认查看的是 filter 表中的规则
```$ iptables -L -v
Chain INPUT (policy ACCEPT 350K packets, 63M bytes)
 pkts bytes target     prot opt in     out     source               destination
Chain FORWARD (policy ACCEPT 0 packets, 0 bytes)
 pkts bytes target     prot opt in     out     source               destination
Chain OUTPUT (policy ACCEPT 18M packets, 1916M bytes)
 pkts bytes target     prot opt in     out     source               destination
```

我们看到三个默认的链，分别是 INPUT、FORWARD 和 OUTPUT，每个链中的第一行输出表示链名称（在本例中为 INPUT/FORWARD/OUTPUT），后跟默认策略（ACCEPT）。

每条链中都可以添加多条规则，规则是按照顺序从前到后执行的。我们来看下规则的表头定义。

- **pkts**：处理过的匹配的报文数量

- **bytes**：累计处理的报文大小（字节数）
- **target**：如果报文与规则匹配，指定目标就会被执行。
- **prot**：协议，例如 `tdp`、`udp`、`icmp` 和 `all`。
- **opt**：很少使用，这一列用于显示 IP 选项。
- **in**：入站网卡。
- **out**：出站网卡。
- **source**：流量的源 IP 地址或子网，或者是 `anywhere`。
- **destination**：流量的目的地 IP 地址或子网，或者是 `anywhere`。

还有一列没有表头，显示在最后，表示规则的选项，作为规则的扩展匹配条件，用来补充前面的几列中的配置。prot、opt、in、out、source 和 destination 和显示在 destination 后面的没有表头的一列扩展条件共同组成匹配规则。当流量匹配这些规则后就会执行 target。

## 总结

以上就是对 iptables 的简要介绍，你已经了解了 iptables 是怎样运行的，规则链及其执行顺序

# Reference

https://jimmysong.io/blog/understanding-iptables/