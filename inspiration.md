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
