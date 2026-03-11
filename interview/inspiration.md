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

## 项目梳理：

### 微服务

1. 起一个微服务需要什么
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
2. 你对微服务的理解 分几个部分
    （1）服务之间的通信
        网关：
            a.提供统一的服务入口，让微服务对前台透明
            b.聚合后台的服务，节省流量，提升性能。
            c.提供安全，过滤，流控等API管理功能
    （2）微服务可能存在的问题。
            a.扩展：扩展软件生命周期开发过程内的任何功能，都可能带来挑战，尤其是在初期。在初始设置期间，重要的是要花时间识别服务之间的依赖关系，并且注意可能破坏向后兼容性的潜在触发因子。到了部署的时候，对自动化的投入至关重要，因为微服务的复杂性使人工部署变得无能为力。 
            b.日志记录：使用分布式系统时，您需要利用集中式日志将所有相关信息集中到一处。否则，积累的日志数量将让您难以招架。
            c.监控：您必须通过一个集中式视图来了解整个系统的情况，以便找出问题的根源。 
            d.调试：无法通过本地集成开发环境（IDE）进行远程调试，因为这种方式无法涵盖数十个或数百个服务。不幸的是，关于应该如何进行调试，目前还没有标准答案。
            e.连接：请考虑使用服务探索功能，无论是集中式的还是集成式
3. k8s 集群 对k8s 的理解 docker swam 转换到k8s
   k8s 集群容器 基本概念了解下
### 日志相关

1.  日志 数数 zap 日志轮转数据丢失（什么情况下产生的）了解日志的轮转机制么
    a.当进行日志轮替的时候，日志文件正在使用中，程序会不停的朝着日志文件写入日志信息，这个时候，直接将日志文件mv重名名进行归档，然后创建一个新文件，重启进程或者重新加载配置文件，那么中间必定有一个中断的时间，这个日志会丢失么？
    lograte 
        a. 延迟重命名法 配置中加入deplaycompress 和 copytruncate 选项，这样lograte 会先复制日志文件，然后清空日志文件，写入新的日志，在下次轮转时在对复制的文件进行压缩
        
    b.数数接入
    c.es日志查询
    
    zap 提供强类型、无反射
    zap 使用零内存分配的json编码器，尽可能避免序列化开销，比其他的结构化日志包快4-10倍
    一个日志库实现需要注意什么　
    elk stack ： 日志落地 搜寻大概流程
    elasticsearch logstash kibana
    filebeat ：轻量的日志采集器
2. 
    将日志内容持久化到文件中，并同时注意磁盘io。上述面试题涉及到的就是这个。
    封装埋点属性

    日志的基本信息需要尽量详细，需要包含文件，函数名，时间等等
    支持不同的日志级别。我们所熟知的DEBUG/INFO/ERROR等等，说的就是这个。
    支持日志切割。支持的维度一般是时间，当然也有根据文件大小的。
    


### 排错
1. 项目故障的定位（查错误日志，看埋点是否出现频率高峰）
   用户反馈  
2. 调试
3. 火焰图
4. jeger tracer opentracing
### 敏感词过滤
1. 讲下敏感词算法（三元tire 树，有限ac 自动机优化？ fail 节点失败后，策略？）
2. 接入阿里云敏感词
### 定时任务
1. 定时服务中的某一个定时任务，使用了一个全局互斥锁，影线业务逻辑处理 
### 活动剧情
### 任务系统
### 邮件系统
1. 邮件策略  
    邮件状态 ： 待发送，已发送，已领取，已过期，撤回
    系统邮件：
    全局邮件：ZSET() score（) 时间戳 
    发送邮件：
        mail_global_filed 全局邮件id集合 
        mail_global_account_field_ + "123456" zset member 是玩家id，成绩是：邮件内容 
        策略检查：账号注册开始时间，注册结束时间；玩家账号等级
        发送邮件记录落地mongo
    领取邮件：
        原子操作：  
        邮件表:标记为已读 
        邮件操作表：inc recv num  
### 玩家数据流向
1. 数据落地
   - 写请求，（是否需要落地到mongo，后面补充上）更新玩家数据，超过30分钟没有写请求后，即使玩家在线，数据也要低频同步到mongo。
   - 心跳低频同步mongo
   - 写请求，玩家不在线（redis中未查找到玩家），mongo拉取玩家数据，更新到redis，
    **数据落地redis时候用事务流水线，减少redis的io次数**
   - 定时任务，过期的缓存数据刷新到db （10s）
2. 数据拉取
   - 路由层 base节点请求，游戏逻辑, FindAccount
3. 缓存数据删除
   - 定时任务，玩家数据落地后，删除玩家缓存数据。 **数据落地redis时候用事务流水线，减少redis的io次数**

### 工具项
1. 时间工具ldate
2. 修复线上数据流程，补发邮件，或者扣除玩家道具数量，可以做到不停服，缩小影响范围
### 排行榜

### 网关
1. 防重放攻击 （恶意攻击，重复领取奖励）
    恶意重复有效的API请求
    时间戳 + 签名 拼接
    上次请求时间戳小于或者等于当前请求时间戳，判定请求非法
    保存当前请求的时间戳，key是玩家的nid

### 支付系统
1. 虚拟服务器  
2. gobot 单元测试  
3. mongo 事务结合 redis  
4. 服务器启动停止  
5. linkcache  
6. 限流器、令牌桶、漏桶、提供一个本地限流器支持、redis 挂掉后 可以维持服务器正常运行  
7. mongo 聚合查询相比mysql、mysql、预加载、索引、引擎啥的了解下    
    操作数组：unwind 
8. 路由层抽象出来
9.  社交服单独做个微服务，支持多开、mongo拉10000个玩家、定时任务会出现内存的峰值
    不从数据拉取优化，从业务上，区分读写好友池，好友业务逻辑不受写好友池的影响 ，要加锁
    好友池 读写分离
10. 中间多层代理，如何 GetLocalIP 获取本机网卡IP
11. 自动部署 CI/CD  讲下流程吧  drone    
12. 优雅启动停止
    ch := make(chan os.Signal,1)
    signal.Notify(ch,syscall.SIGTERM,syscall.SIGINT)
    for {
        select {
            case : syscall.SIGTERM
            case :
        }
    }
13. 心跳
    矫正客户端时间
    体力恢复时间
    维护开始前10分钟提示  心跳 + 主动通知 + 状态改变
    玩家数据低频强制落地（全量）


14. 排行榜
    表结构

15. 社交服务
    基本结构：好友属性，玩家服属性。
    好友池：拉取好友列表，好友池，好友池数据落地，好友池数据拉取，好友池数据删除
    好友池数据落地：定时任务，玩家数据落地后，删除玩家缓存数据。 **数据落地redis时候用事务流水线，减少redis的io次数**，落地mongo时候用事务
    好友池数据拉取：玩家上线时候，拉取好友池数据，拉取玩家数据，拉取玩家好友列表
    好友池数据删除：玩家下线时候，删除好友池数据，删除玩家缓存数据，删除玩家数据
    
支付

### mongo

## 聚合
1. 数组字段查询
     
## 事务
ssl/TSL
实际应用中使用对称加密 

3. 分组
    group 
    {
        $group : 
            {
                _id:"xxx", 
                count :{ $sum : 1}
            }
    }
    {
        $match : 
        {
            count : {$gt :1 }
        }
    }
    