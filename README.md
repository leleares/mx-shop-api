# mx-shop-api
调用底层grpc服务暴露为上层http服务。

# user-web 目录
负责暴露底层user的grpc服务为上层http服务。

# go日志库 zap
分为logger和sugarLogger，sugarLogger提供简单易用的日志打印api，logger打印日志api用起来稍复杂但是性能极致。
日志是分级别的，例如分开发环境、生产环境。
debug、info、warn、error、fetal。
zap.L是zap.Logger的简易调用方式，zap.S是zap.SugaredLogger的简易调用方式，前者性能更好但需明确说明数据类型，后者调用更方便。

# 使用 protoc 生成 go 代码
生成普通proto结构体代码: `protoc --go_out=. user.proto`
生成gRPC service接口代码:  `protoc --go-grpc_out=. user.proto`

# DTO
DTO（Data Transfer Object）

# go的配置文件管理
viper
why viper? 支持默认值、监听配置文件变动、很多简单易用的能力。

# redis
基于内存的 Key-Value 数据库

启动 redis：`brew services start redis`
测试 redis 是否运行成功：`redis-cli ping`
启动redis服务端：`redis-server`
启动redis客户端：`redis-cli`

| 配置项          | 值               |
| ------------ | --------------- |
| **Host**     | `127.0.0.1`     |
| **Port**     | `6379`          |
| **Password** | 空（如果你没设置密码的话）   |

# 服务注册 服务发现
启动 consul：`consul agent -dev`
访问可视化界面：`http://localhost:8500/ui`
使用dig解析服务name对应ip和port：`dig @127.0.0.1 -p 8600 web.service.consul` 其中：dig @127.0.0.1 -p 8600 表示连接本地的consul服务，解析web服务。xxx.service.consul中xxx表示服务名称，后面为固定写法。

# 常用的负载均衡算法
1. 轮询法（Round Robin） （平均将请求分配给各个服务器）
2. 随机法，同一
3. 源地址哈希法（大致意思是通过某种算法，使得同一个客户端IP访问的始终是同一台服务器）
4. 加权轮询（考虑机器性能等情况）
5. 加权随机（考虑机器性能等情况）
6. 最小连接数（考虑服务器的连接数，将请求分配给连接数较小的服务器）

# 分布式配置中心选型
apollo: 携程开源，大而全
nacos: 阿里开源，小而全

### nacos
Nacos 是一个 Java 服务，本质是一个 Web 应用
本地启动nacos：进入到dev目录 进入nacos目录执行：`sh bin/startup.sh -m standalone`
访问：`http://127.0.0.1:8848/nacos`

nacos中的一些概念：
- 命名空间 - 可以理解为一个项目就可以创建一个命令空间，例如user-web一个命名空间，user-srv一个命名空间。
- 组 - 可以用来做环境隔离，例如dev组，prod组。
- 配置集（data id）可以理解为具体的配置文件。

### ngrok
内网穿透工具，可将本地指定端口服务映射为公网IP。
用法: `ngrok http xxxx` xxxx 为本地服务端口。

### *bool问题
```go
type Form struct {
	Checked *bool `json:"checked" form:"checked" binding:"required"`
}

// bool为什么要整成指针类型，bool类型会有什么问题？
// 当设置为bool类型时，客户端传false进来，gin会认为这是bool的零值，会忽略，从而认为没传checked字段，正好和required冲突。
// 设置为bool指针，会认为nil才是零值，参数正常传递true/false，但使用时注意*Checked取到bool值。
```

### 沙箱环境
一般用于给开发测试阶段提供的完整功能（无需上传相关证书以及证明），等上线后更改配置即可，例如对于支付宝提供的sandbox，无需提供任何信息，可以测试验证支付功能，模拟支付，不会真的扣钱。

### 对称加密与非对称加密
1. 对称加密：只有一个密钥，安全性相对较低。
2. 非对称加密：有公钥和私钥之分，公钥可以交给别人，私钥需要自己保护好。私钥加密使用公钥才能解开。

### elasticsearch
mysql实现搜索的痛点：
1. 性能低
2. 没有相关度排名
3. 无法全文搜索
4. 没有分词

什么是elasticsearch？
Elasticsearch 是一个分布式可扩展的实时搜索和分析引擎，一个建立在全文搜索引擎 Apache Lucene（TM）基础上的搜索引擎：
1. 分布式实时文件存储，并将每一个字段都编入索引，使其可以被搜索。
2. 实时分析的分布式搜索引擎。
3. 可以扩展到上百台服务器，处理PB级别的结构化或非结构化数据。I


### 
```bash
# 使用docker启动es
docker run -d \
  --name es \
  -p 9200:9200 \
  -e "discovery.type=single-node" \
  -e "ES_JAVA_OPTS=-Xms512m -Xmx512m" \
  -e "xpack.security.enabled=false" \
  -v /Users/$(whoami)/es-data:/usr/share/elasticsearch/data \
  -v /Users/$(whoami)/ik:/usr/share/elasticsearch/config/analysis-ik \
  docker.elastic.co/elasticsearch/elasticsearch:8.11.0

# 启动kibana
docker run -d \
  --name kibana \
  -p 5601:5601 \
  --link es:elasticsearch \
  docker.elastic.co/kibana/kibana:8.11.0  
```

### es容器相关操作
进入es容器：`docker exec -it es /bin/bash`
以root身份进入es容器`docker exec -it -u root es /bin/bash`


### 启动rocketmq
根目录下有一个docker-compose文件，直接进入根目录执行命令即可
> 启动rocketmq `docker-compose up -d`
> 关闭rocketmq `docker-compose down`
>rocketmq webUI：`http://localhost:8080/`

### 启动jaeger
```bash
docker run --rm --name jaeger \
  -p 16686:16686 \
  -p 6831:6831/udp \
  -p 14268:14268 \
  jaegertracing/all-in-one:1.57
```

> jarger webUI：`http://localhost:16686/`
>
> ### 启动kong

`cd kong`
`docker compose up -d`

### cat写入文件操作
```bash
# 覆盖写入
cat > custom.dic <<EOF
iphone15
华为mate60
小米14
EOF

# 追加写入
cat >> custom.dic <<EOF
特斯拉
拼多多
EOF

# 单行追加
echo "macbookpro" >> custom.dic

echo "mate40" >> custom.dic
```

### jenkins 
jenkins 做的事情本质上是自动构建部署。
jenkins 是一个服务器，可以通过任务的方式来执行一项构建部署动作。这个任务可以指定关联仓库，指定什么时候pull code，指定构建触发器（什么行为触发项目构建（build 项目））、也可指定构建完成后做什么事情（例如可以指定将打包后的文件上传到其他服务器，就等于是部署了）。
创建任务有freestyle和pipeline两种模式，freestyle就是以UI方式指定一些行为进行构建部署，pipeline就是将前面的配置固化成特定脚本(Jenkinsfile)，让jenkins服务执行即可创建任务整个构建部署动作。

自动构建部署：CI/CD 自动化。CI：开发者将代码合到develop，推到远程分支，自动触发构建。CD：构建通过后自动发布到测试环境。
Pipeline 不只是固化脚本，更核心的思想是：流水线即代码，Jenkinsfile 是可以跟着代码一起放进仓库的，能版本控制、能review、能回滚。
构建触发器的类型：定时构建（指定某个时间则进行构建动作）、轮询构建（几分钟构建一次）、webhook构建（git push通知jenkins来进行构建）、其他任务构建触发本任务构建。

### 项目构建部署
实际总共有三层
git 仓库 -> 构建服务器 -> 运行服务器
现在完整的电商项目有这么几个服务：
git仓库：
前端服务（vue）
后端服务（go和python）

前端服务构建流程：构建服务器拉取前端vue代码 -> 由于构建服务器本地前置已安装了nodejs环境，因此通过 npm i -f 安装依赖， npm run build 打包生成 dist 静态目录 -> 构建服务器将其上传到 nginx 即可。  （运行服务器无需配置nodejs环境，因为？？？）
后端服务（go）构建流程：构建服务器拉取go代码 -> 由于构建服务器本地前置已安装了go环境，因此通过 go get 安装依赖，go build 生成可执行文件 -> 构建服务器将其上传到运行服务器可直接运行go服务。（运行服务器无需配置go环境，因为拿来的就是可执行文件） 
后端服务（python）构建流程：构建服务器拉取python代码 -> 构建服务器一般将其源码上传到运行服务器 -> 因此需要运行服务器安装python运行环境，执行python源码运行服务。

### 前端渲染过程
将打包后的dist目录上传到服务器 -> 用户访问域名 https.xxx.com/ -> DNS 域名解析找到服务器 IP，服务器响应给浏览器 html 文件 -> 浏览器解析执行 html -> 发现引用了css，则请求 css 文件。引用了 js 则请求 js 文件....，最后执行js，整个前端流程渲染完毕 -> 页面交互再去请求服务器数据做渲染。

### go 项目 CI/CD 全流程
改完代码 push 到 master 分支 -> jenkins 任务监测到 master 分支变化，触发构建 -> go build 打包成二进制可执行文件 -> 上传到另一台运行服务器并且指定启动命令（一般通过 ssh 来达到启停运行服务器中该 go 进程）

问题1: go 进程在运行服务器中如何进行管理，例如重启、保活等操作？
通过：systemd （可以理解为pm2，进程管理工具而已）

### k8s是啥？ 可以取代jenkins?
不能，jenkins 和 k8s 是相互配合的，k8s 负责进程、容器、机器管理调度相关工作。
工作流程：改完代码 push 到 master 分支 -> jenkins 任务监测到 master 分支变化，触发构建 -> 打包成 docker 可运行镜像 -> 推到git仓库 -> 触发k8s（触发方式可以是jenkins触发，也可以是一些gitops） -> k8s负责运行镜像到什么容器什么机器甚至负责整个集群调度。

### 服务端项目迭代
单体应用 -> 微服务 -> 服务网格 (service mesh) -> 云原生 (serverless)

总体趋势就是：尽量让业务代码保持纯粹，将服务治理（负载均衡、服务注册发现、限流熔断等）从业务代码中分离。逐级把"非业务复杂度"外包出去。

服务网格：将微服务中所有服务治理代码下沉到一个基础设施中，陪在微服务旁边叫做边车代理，替微服务完成服务治理。

云原生：除了写业务函数以外，无需关心任何非业务相关的内容，所有服务都是在云上自治理。

