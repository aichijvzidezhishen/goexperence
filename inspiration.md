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


### 项目梳理：
braid 项目脚手架 负责节点调度 服务治理 grpc 通信() 消息队列
1. 日志 数数 zap 日志轮转数据丢失（什么情况下产生的）了解日志的轮转机制么
2. 起一个微服务需要什么
3. 你对微服务的理解 分几个部分
4. k8s 集群 对k8s 的理解 docker swam 转换到k8s
5. 项目故障的定位（查错误日志）
6. 讲下敏感词算法（三元tire 树，有限ac 自动机优化？ fail 节点失败后，策略？）
7. 定时服务中的某一个定时任务，使用了一个全局互斥锁，影线业务逻辑处理 
8. 修复线上数据流程，补发邮件，或者扣除玩家道具数量，可以做到不停服，缩小影响范围
9. 跨槽事务，精通redis集群原理支持（一致性hash 算法） 主从复制 哨兵模式？ 
10. 消息队列 krafa、redis 、nsp（自己实现下）
11. 邮件策略
12. 虚拟服务器
13. gobot  单元测试 
14. mongo 事务结合 redis 
15. 服务器启动停止
16. linkcache
17. 限流器 令牌桶、漏桶、提供一个本地支持、redis 挂掉后 可以维持服务器正常运行
18. redis 数据结构 zset、set、hash、string、
  zset:1/(1-q)
19. mongo 聚合查询 相比mysql 、 mysql 预加载、索引、引擎啥的了解下
20. 路由层抽象出来
21. 社交服单独做个微服务，支持多开、mongo拉10000个玩家、定时任务会出现内存的峰值
    不从数据拉取优化，从业务上，区分读写好友池， 好友业务逻辑不受写好友池的影响 ，要加锁






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
