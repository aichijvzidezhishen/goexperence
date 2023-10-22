# goexperence
## 算法与数据结构
### 
代理

## hw/base
1. 无缓冲 Chan 的发送和接收是否同步
2. go语言的并发机制以及它所使用的CSP并发模型
3. Golang 中常用的并发模型？
4. JSON 标准库对 nil slice 和 空 slice 的处理是一致的吗？
5. 闭包
6. 协程，线程，进程的区别
7. 互斥锁，读写锁，死锁问题是怎么解决(场景)
8. Golang的内存模型（小对象多了会增加内存压力）
9. slice传参有几种方式不改变原始数据
10. epoll 、poll、select
11. channel线程安全
12. map 线程安全
13. golang gc
14. goroutine 调度
15. 并发和并行
16. uintptr 和 unsafe.Pointer 的区别
    uintptr 和 unsafe.Pointer 是 Go 语言中用于处理指针的类型，它们的使用场景和特性有所不同。

    uintptr 是一个无符号整数类型，用于存储指针的数值。它可以用于将指针转换为整数，以及将整数转换回指针。但是，uintptr 并不保留指针的类型信息，因此在将整数转换回指针时需要确保类型的正确性。uintptr 主要用于底层编程，例如在进行指针运算或与 C 语言进行交互时。

    unsafe.Pointer 是一个特殊的指针类型，可以用于存储任意类型的指针。它可以将指针转换为 unsafe.Pointer 类型，以及将 unsafe.Pointer 转换回原始指针类型。unsafe.Pointer 的使用需要谨慎，因为它可以绕过 Go 语言的类型系统，可能导致类型不匹配或内存安全问题。因此，它主要用于与 C 语言进行交互或进行底层编程时，例如在使用 cgo 包调用 C 函数时。

    总结来说，uintptr 用于将指针转换为整数，而 unsafe.Pointer 用于在指针和不同类型之间进行转换。在一般情况下，应尽量避免使用 uintptr 和 unsafe.Pointer，而是使用 Go 语言提供的类型安全的指针操作和类型转换机制。只有在必要的情况下，才应该使用 uintptr 和 unsafe.Pointer，并且需要谨慎处理类型和内存安全性。

17. 反向代理- 负载均衡器
19. goroutinue channel mutex 应用场景
20. log 线程安全么  
21. Golang 里的逃逸分析是什么？怎么避免内存逃逸？
22. 配置中心如何保证一致性？
23. Golang 的GC触发时机是什么?
24. Redis 里数据结构的实现熟悉吗?
25. Etcd的Raft一致性算法原理?

## 微服务架构
1. swarm选举过程
2. 关于任期
3. 关于随机超时
4. SLB原理
5. 分布式一致性原则
6. 如何保证服务宕机造成的分布式服务节点处理问题?
7. 服务发现怎么实现的

## 数据库
### mysql
1. 索引优化，如何查询优化结果
2. 联合索引 怎么写？ 规则，生效时间
3. 索引原理 b树

### mongodb
1. 事务 session   
2. 文档数据库 高性能、高可用、易扩展
3. mongo 分片 解决系统增长：垂直拓展、水平拓展
    mongo的分片集群包含以下组件：


### redis 
1. 事务性
2. 数据持久化方式
3. 缓存淘汰策略
4. [分布式锁实现](hw\redis-lock\lock.go)
5. Pub/Sub模型
6. Redis的线程模型

### PostgreSQL
1. MVCC机制
2. 索引优化
3. 事务隔离级别
4. 触发器和存储过程
5. 分区表

### Distributed Systems
1. CAP理论
2. 一致性哈希算法
3. Paxos和Raft算法
4. 分布式事务
5. 分布式锁
6. 服务发现和注册
7. 负载均衡策略
8. 分布式存储系统


## NetWork
1. 公网 内网 专用网络
2. 代理
3. websocket 属于网络分层那一层
4. https 流程是怎样的（公私钥、数字证书、加密、对称加密、非对称加密。）
5. 说说HTTP的状态码，301和302的区别？
6. 对称加密和非对称加密  
7. dsa 
8. dns解析过程
9. csrf 攻击 如何避免
10. websocket 和 socket 区别
11. DoS、DDoS、DRDoS攻击
12. xxs攻击
13. http请求的过程和原理
14. forward 和redirect 的区别
    直接转发方式 
    间接转发方式
15. sql 注入，如何攻击
16. 如何预防sql 注入问题
17. Session和Cookie的区别。
18. IP地址有哪些分类？
19. ARP 协议的工作过程
20. TCP 和 UDP 分别对应的常见应用层协议
21. 保活计时器
22. 服务器出现了大量CLOSE_WAIT状态如何解决
23.  URI和URL的区别
24.  ICMP协议的功能
25.  ping的原理
26.  TCP 的三次握手机制

    [![三次握手]](http://mianbaoban-assets.oss-cn-shenzhen.aliyuncs.com/xinyu-images/MBXY-CR-f795189dd1254b2b5765b0a617198b9c.png)
27. TCP四次挥手过程
28. TCP的粘包和拆包
29. 说说半连接队列和 SYN Flood攻击的关系
30. TCP的滑动窗口
31. TCP的拥塞控制
32. 慢启动算法
33. 拥塞避免算法
34. TCP报文首部有哪些字段，作用
    1. 16位端口号：源端口号，主机该报文段是来自哪里；目标端口号，要传给哪个上层协议或应用程序

    2. 32位序号：一次TCP通信（从TCP连接建立到断开）过程中某一个传输方向上的字节流的每个字节的编号。
    3. 32位确认号：用作对另一方发送的tcp报文段的响应。其值是收到的TCP报文段的序号值加1。
    4. 4位头部长度：表示tcp头部有多少个32bit字（4字节）。因为4位最大能标识15，所以TCP头部最长是60字节。
    5. 6位标志位：URG(紧急指针是否有效)，ACk（表示确认号是否有效），PSH（缓冲区尚未填满），RST（表示要求对方重新建立连接），SYN（建立连接消息标志接），FIN（表示告知对方本端要关闭连接了）
    6. 16位窗口大小：是TCP流量控制的一个手段。这里说的窗口，指的是接收通告窗口。它告诉对方本端的TCP接收缓冲区还能容纳多少字节的数据，这样对方就可以控制发送数据的速度。
    7. 16位校验和：由发送端填充，接收端对TCP报文段执行CRC算法以检验TCP报文段在传输过程中是否损坏。注意，这个校验不仅包括TCP头部，也包括数据部分。这也是TCP可靠传输的一个重要保障。
    8. 16位紧急指针：一个正的偏移量。它和序号字段的值相加表示最后一个紧急数据的下一字节的序号。因此，确切地说，这个字段是紧急指针相对当前序号的偏移，不妨称之为紧急偏移。TCP的紧急指针是发送端向接收端发送紧急数据的方法。
35. 超时重传
36. 重复SACK（D-SACK）
37. http 常见的端口
    - 21:FTP（文件传输协议）
    - 22：ssh
    - 25:SMTP（简单邮件传输协议）
    - 53:DNS 域名服务器
    - 80:HTTP 超文件传输协议
    - 110：POP3 邮件协议3
    - 443:HTTPS
    - 1080:Sockets
    - 1521：Oracle 数据库默认端口
    - 3306：MYSQL 服务默认端口

## 操作系统
### 操作系统结构
1. 内核
2. 用户态和内核态
3. 用户态和内核态是如何切换

### 进程和线程
1. 并行和并发有什么区别
2. 进程上下文切换
    进程既可以在用户空间运行，又可以在内核空间运行。
    - 前后两个线程属于不同进程。此时，因为资源不共享，所以切换过程就跟进程上下文切换是一样。
    - 前后两个线程属于同一个进程。此时，因为虚拟内存是共享的，所以在切换时，虚拟内存这些资源就保持不动，只需要切换线程的私有数据、寄存器等不共享的数据
3. 进程有哪些状态
4. 僵尸进程
5. 孤儿进程
6. 进程有哪些调度算法
7. 进程间通信有哪些方式
8. 进程和线程的联系和区别
   - 线程和进程的联系
   - 线程与进程的⽐较如下

9. 线程有哪些实现方式
    - 内核态线程实现
    - ⽤户态线程实现
    - 混合线程实现

10. 线程间如何同步
11. 死锁
12. 死锁产生有哪些条件？
13. 避免死锁
14. 活锁和饥饿锁

### 虚拟内存
1. 定义
2. 内存分段
3. 内存分页
4. 多级页表
5. 块表
6. 分页和分段
7. 交换空间
8. 页面置换算法
9. 硬连接和软连接的区别

### IO
1. 零拷贝
2. 阻塞与⾮阻塞 **I/O **、 同步与异步 I/O？
3. IO 多路复用