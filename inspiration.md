# 灵感
## nginx 正反向代理  
1. 正向代理与反向代理的区别  
    所谓代理服务器就是位于发起请求的客户端与原始服务器端之间的一台跳板服务器，正向代理可以隐藏客户端，反向代理可以隐藏原始服务器。  
    1. 正向代理的概念  
    2. 反向代理的概念

    
   
## 本地虚拟集群环境 docker-swarm 部署
>https://kiwenlau.com/2016/07/03/vagrant-vm-cluster/

## redis + lua 实现分布式 限流器

## go工作区模式
> Go工作区模式是Go开发者Michael Matloob在2021年4月提出的一个名为“Multi-Module Workspaces in cmd/go”的proposal。这个proposal引入一个go.work文件用于开启Go工作区模式。go.work通过directory指示符设置一些本地路径，这些路径下的go module构成一个工作区(workspace)，Go命令可以操作这些路径下的go module，也会优先使用工作区中的go module。
1.   
        
        ```
        go install golang.org/dl/gotip@latest 
        // go 1.17版本及以后使用go install。go 1.16及之前的版本用go get
        $gotip download
        $gotip version
        go version devel go1.18-b7529c3 Tue Nov 9 06:27:04 2021 +0000 darwin/amd64
        ```
## 配置命令行代理
export https_proxy=http://127.0.0.1:7890 http_proxy=http://127.0.0.1:7890 all_proxy=socks5://127.0.0.1:7890

redis 事务
## 同一事务中操作一个key，会有跨槽风险

当一个命令在事务执行，失败后，redis会中止事务的执行，并且不会执行官事务队列中剩余的命令，已经成功执行的命令不会撤销或者回滚。 

## 碰到一个切片浅拷贝的问题

··· go
        type Event struct {
        Eventid int32
        Idlist  []string
    }

    type Events struct {
        List []Event
    }

    func (es *Events) Refresh() {
        for _, v := range es.List {
            v.RefreshEvent()
        }
    }

    func (e *Event) RefreshEvent() {
        PrintSliceStruct(&e.Idlist, "RefreshEvent")
        for k := range e.Idlist {
            e.Idlist[k] = "test" + strconv.Itoa(int(e.Eventid))
        }
    }

    func PrintSliceStruct(s1 *[]string, from string) {
        sh := (*reflect.SliceHeader)(unsafe.Pointer(s1))
        fmt.Printf("from %v slice addr %+v\n", from, sh)
    }

    type Slice struct {
        unsafe.Pointer
        len 
        cap 
    }
    浅拷贝 复制出来的对象和原对象指向同一地址
    var a []int32{1,2,3}
    //b 会新开辟一块内同空间，指向新的地址，但是b的底层数组指针还是和原切片相同
    b := a
    b1 := a[:]
    b2 := a[start:end]

    //深拷贝
    copy() 内置函数
    a := []int{1, 2, 3}
    b := []int{-1, -2, -3, -4}
    copy(b, a)
    fmt.Println(unsafe.Pointer(&a))  // 0xc0000a4018
    fmt.Println(a, &a[0])            // [1 2 3] 0xc0000b4000
    fmt.Println(unsafe.Pointer(&b))  // 0xc0000a4030
    fmt.Println(b, &b[0])    // [1,2,3,-4]
···

### 项目梳理：
braid 项目脚手架 负责节点调度 服务治理 grpc 通信() 消息队列
1. 日志 数数 zap 日志轮转数据丢失（什么情况下产生的）了解日志的轮转机制么
    当进行日志轮替的时候，日志文件正在使用中，程序会不停的朝着日志文件写入日志信息，这个时候，直接将日志文件mv重名名进行归档，然后创建一个新文件，重启进程或者重新加载配置文件，那么中间必定有一个中断的时间，这个日志会丢失么？


2. 起一个微服务需要什么
    一致性 加权轮询（每个节点存储的数据相同）
    hash算法扩容？？
    一致性哈希解决了扩容缩容问题  对2^32取模 哈希环
        a. 对存储节点哈希
        b.  对key哈希 首先，对 key 进行哈希计算，确定此 key 在环上的位置；
            然后，从这个位置沿着顺时针方向走，遇到的第一节点就是存储 key 的节点。
            因此，在一致哈希算法中，如果增加或者移除一个节点，仅影响该节点在哈希环上顺时针相邻的后继节点，其它数据也不会受到影响。
            但是一致性哈希算法并不保证节点能够在哈希环上分布均匀，这样就会带来一个问题，会有大量的请求集中在一个节点上。
        c.通过虚拟节点提高均衡度
          不再将真实节点映射到哈希环上，而是将虚拟节点映射到哈希环上，并将虚拟节点映射到实际节点，所以这里有「两层」映射关系。
3. 你对微服务的理解 分几个部分
    （1）服务之间的通信
        网关：
            a.提供统一的服务入口，让微服务对前台透明
            b.聚合后台的服务，节省流量，提升性能。
            c.提供安全，过滤，流控等API管理功能
    （2）微服务可能存在的问题。·
            a.扩展：扩展软件生命周期开发过程内的任何功能，都可能带来挑战，尤其是在初期。在初始设置期间，重要的是要花时间识别服务之间的依赖关系，并且注意可能破坏向后兼容性的潜在触发因子。到了部署的时候，对自动化的投入至关重要，因为微服务的复杂性使人工部署变得无能为力。 
            b.日志记录：使用分布式系统时，您需要利用集中式日志将所有相关信息集中到一处。否则，积累的日志数量将让您难以招架。
            c.监控：您必须通过一个集中式视图来了解整个系统的情况，以便找出问题的根源。 
            d.调试：无法通过本地集成开发环境（IDE）进行远程调试，因为这种方式无法涵盖数十个或数百个服务。不幸的是，关于应该如何进行调试，目前还没有标准答案。
            e.连接：请考虑使用服务探索功能，无论是集中式的还是集成式
4. k8s 集群 对k8s 的理解 docker swam 转换到k8s
5. 项目故障的定位（查错误日志）
6. 讲下敏感词算法（三元tire 树，有限ac 自动机优化？ fail 节点失败后，策略？）
7. 定时服务中的某一个定时任务，使用了一个全局互斥锁，影线业务逻辑处理 
8. 修复线上数据流程，补发邮件，或者扣除玩家道具数量，可以做到不停服，缩小影响范围
9.  跨槽事务，精通redis集群原理支持（一致性hash 算法） 主从复制 哨兵模式？ 
10. 消息队列 krafa、redis 、nsp（自己实现下）
11. 邮件策略
12. 虚拟服务器
13. gobot  单元测试 
14. mongo 事务结合 redis 
    
15. 服务器启动停止
16. linkcache
17. 限流器 令牌桶、漏桶、提供一个本地限流器支持、redis 挂掉后 可以维持服务器正常运行
18. redis 数据结构 zset、set、hash、string
19. mongo 聚合查询 相比mysql 、 mysql 预加载、索引、引擎啥的了解下
    操作数组：unwind
20. 路由层抽象出来
21. 社交服单独做个微服务，支持多开、mongo拉10000个玩家、定时任务会出现内存的峰值
    不从数据拉取优化，从业务上，区分读写好友池， 好友业务逻辑不受写好友池的影响 ，要加锁
22. 中间多层代理，如何 GetLocalIP 获取本机网卡IP
23. Aoi 十字链表，灯塔，九宫格
    a. 插入  链表长度很长时，要从头开始遍历，效率低（可以用快排，分治，按范围划分）
    b. 删除  链表长度很长时，要从头开始遍历，效率低（可以用快排，分治，按范围


### mongo

## 聚合
1.  $project filter cond  

    $filter: {
                    input: "$baginfo.items",
                    as: "item",
                    cond: {
                        $in: ["$$item.dictid", {
                            $ifNull: ["$baseInfo.usemedallist", []]
                        }]
                    }
                }
2. 数组字段查询
    a. 过滤数组中不需要的字段 $not $in
    {{"items.id:  {$not :{$in [1,2,3]}}}}
    b. $unwind
    c. push 
    
3. 分组
    group 
    {$group : {_id:"xxx",count :{ $sum : 1}}}
    {$match : {count : {$gt :1 }}}
