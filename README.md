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
5. 协程，线程，进程的区别
6. 互斥锁，读写锁，死锁问题是怎么解决(场景)
    
7. Golang的内存模型（小对象多了会增加内存压力）
8. slice传参有几种方式不改变原始数据
    

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

23. 微服务概念.  
选举，负载，服务器优雅启停 

1.  SLB原理.

2.  分布式一直性原则.

3.  如何保证服务宕机造成的分布式服务节点处理问题?

4.  服务发现怎么实现的.


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
4. 
