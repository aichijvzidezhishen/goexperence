# goexperence

**增dasd我**
>都说了就拉上来看v才能到家啊是的,非常我的是否我的 我的世界，


--string
## 算法与数据结构
### 
代理

## hw/base
1. 无缓冲 Chan 的发送和接收是否同步  
ch := make(chan int)    `无缓冲的channel`由于没有缓冲发送和接收需要同步.  
ch := make(chan int, 2) `有缓冲channel`不要求发送和接收操作同步.   

---
- channel无缓冲时，发送阻塞直到数据被接收，接收阻塞直到读到数据。
- channel有缓冲时，当缓冲满时发送阻塞，当缓冲空时接收阻塞。  
2. go语言的并发机制以及它所使用的CSP并发模型
3. Golang 中常用的并发模型？
4. JSON 标准库对 nil slice 和 空 slice 的处理是一致的吗？
5. 闭包
6. 协程，线程，进程的区别
7. 互斥锁，读写锁，死锁问题是怎么解决(场景)
    
8. Golang的内存模型（小对象多了会增加内存压力）
9. slice传参有几种方式不改变原始数据
    

    传入数组 **尽量不使用slice指针传参**
    ```
    var a = []int{2,3,5}
    
    func test (a [3]int) {

    }
    ``` 
9.  epoll 、poll、select

10. channel线程安全
11. map 线程安全
12. golang gc
13. goroutine 调度
14. 并发和并行
15. uintptr 和 unsafe.Pointer 的区别

16. 反向代理- 负载均衡器

17. 分布式锁实现
    - redis + lua脚本、memcached、tair
    - 基于数据库实现分布式锁
    - 基于Zookeeper实现分布式锁
17. goroutinue channel mutex 应用场景
    mutex 用于顺序的访问资源
    channel 用于在 goroutine 之间编排计算。
    数据流动
19. log 线程安全么  
    Golang的标准库提供了log的机制，但是该模块的功能较为简单（看似简单，其实他有他的设计思路）。在输出的位置做了线程安全的保护
20. Golang 里的逃逸分析是什么？怎么避免内存逃逸？

21. 配置中心如何保证一致性？

22. Golang 的GC触发时机是什么?
    

Redis 里数据结构的实现熟悉吗?

22. Etcd的Raft一致性算法原理?


## 微服务架构
  微服务概念.  
    选举，负载，服务器优雅启停 
    1. swarm选举过程
        - 初始状态
            Raft 节点初始化心跳时间都是随机的  任期编号都是0 角色都是跟随者
        - 请求投票
            此时没有一个节点是领导者，节点等待超时心跳后，会推荐自己成为候选人，后向集群其他节点发起投票信息，此时任期编号+1，自荐会获取
            自己的一票选举

        - 跟随者投票
            跟随者收到投票信息后，如果候选人符合投票要求后，则将自己宝贵的一票（任期内跟随者只能投给先来的候选人一票，后来的候选人不能投票给他了）投给该候选人，同时更新任期编号。

        - 当选领导者
            当节点c获得大多数投票后 他会成为本期任期内的领导者
        - 领导者与跟随者保持心跳
            领导者周期性发送心跳给周围节点，告知自己是领导，同时刷新跟随者的超时时间，防止跟随者发起新的领导者选举。
    2. 关于任期

    1. 关于随机超时
    用于发现和配置服务的工具。提供了可供客户端注册和发现服务的API。consul可以对
    服务进行健康检查，以确定服务的可用性。 key val 键值对存储
    jaeger 服务追踪

    
1.  SLB原理.

2.  分布式一直性原则.

3.  如何保证服务宕机造成的分布式服务节点处理问题?

4.  服务发现怎么实现的.
    基于现代云微服务应用，服务实例具有动态分配的网络位置，由于动态扩索，故障与升级，整组服务实例会动态变更。
    客户端发现
    服务端发现

## 数据库
### mysql
1. 索引优化，如何查询优化结果
2. 联合索引 怎么写？ 规则，生效时间
3. 索引原理 b树

### mongodb
1. 事务 session   
    a. dasd

2. 
3. 
## NetWork
1. 公网 内网 专用网络
2. 代理
3. websocket 属于网络分层那一层
4. https 流程是怎样的（公私钥、数字证书、加密、对称加密、非对称加密。）
5. 说说HTTP的状态码，301和302的区别？
    + HTTP状态码，302和301都有重定向的含义，但是它们也是有区别的。

    + 301：（永久性转移）请求的网页已被永久移动到新位置。服务器返回此响应时，会自动将请求者转到新位置。

    + 302：（暂时性转移）服务器目前正从不同位置的网页响应请求，但请求者应继续使用原有位置来进行以后的请求。此代码与响应GET和HEAD请求的301代码类似，会自动将请求者转到不同的位置。

6. 对称加密和非对称加密  
    - 非对称加密，加密速度慢，会造成等待时间过长的问题.

    - 非对称加密结合对称加密的方式：非对称加密传输所需的密钥（只有第一次传输密钥用非对称加密方式），后面的通信使用对称加密的方式加密解密.
    
    - 我们无法确定得到的非对称加密里的公钥一定是安全的公钥，可能存在中间人截取对方发给我们的公钥，替换成自己的公钥发送给我们，当我们使用他的公钥加密后发送的信息，可能被他用自己的私钥解密.  他可以伪装成我们以同样的方法向对方发送消息，这样我们的信息就会被窃取

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
4. 

### 进程和线程
1. 并行和并发有什么区别
2. 进程上下文切换
3. 进程有哪些状态
4. 僵尸进程
5. 孤儿进程
6. 进程有哪些调度算法
   - **先来先服务**
   - **短作业优先**
   - 优先级调度
   - 时间片轮转
   - 最短剩余时间优先
7. 进程间通信有哪些方式
   - 管道
   - 信号
   - 消息队列
   - 共享内存
   - 信号量
   - 优缺点
        - 管道：简单；效率低，容量有限；
        - 消息队列：不及时，写入和读取需要用户态、内核态拷贝。
        - 共享内存区：能够很容易控制容量，速度快，但需要注意不同进程的同步问题。
        - 信号量：不能传递复杂消息，一般用来实现进程间的同步；
        - 信号：它是进程间通信的唯一异步机制。
        - Socket：用于不同主机进程间的通信。
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

