## happylemon
### 预览
```
/code/go/src/happylemon$./happylemon

####################################################
./happylemon -mode=cmd (-conf=conf/wcd.toml) -cmd=sync -method=testLog 命令行模式
./happylemon -mode=http (-conf=conf/wcd.toml) http服务模式
####################################################
```
http服务模式
```
/code/go/src/happylemon$./happylemon -mode=http
[GIN-debug] GET    /db                       --> happylemon/controllers.(*TestController).TestDb-fm (3 handlers)
[GIN-debug] GET    /redis                    --> happylemon/controllers.(*TestController).TestRedis-fm (3 handlers)
[GIN-debug] POST   /post                     --> happylemon/controllers.(*TestController).PostData-fm (3 handlers)
[GIN-debug] GET    /testapigrpc              --> happylemon/controllers.(*TestController).ApiGrpc-fm (3 handlers)
[GIN-debug] GET    /httpstatus               --> happylemon/controllers.(*TestController).Status-fm (3 handlers)
[GIN-debug] GET    /testlog                  --> happylemon/controllers.(*TestController).WcdLog-fm (2 handlers)
[GIN-debug] GET    /metrics                  --> happylemon/lib/prometheus.Handler.func1 (3 handlers)
Server Port :8888
[GIN-debug] Listening and serving HTTP on :8888
```
命令行模式
```
./happylemon -mode=cmd -cmd=test
```
### 安装
```
我的golang版本
/code/go/src/happylemon$go version
go version go1.10.3 darwin/amd64
go clone github.com/weichangdong/happylemon 或者直接下载zip包
```
需要的第三方库都在vendor里面了,不需要额外go get.

### 基本说明
1. 使用gin框架构建的一个mvc(controllers,models,entity,service)模式框架,主要是提供api和命令行模式(后台异步执行一些逻辑)执行命令.可通过参数指定运行模式.基于之前开发的一个项目改的,相比之前的项目,做了几点优化.
2. 之前命令行模式需要单独go build,然后提供api接口的需要go build一个,等于是两个二进制文件,这样面临的问题有
	1. 命令行执行某个操作完成之后通过channel通知api,但是因为不是一个应用,channel没法通信.
	2. 代码发布的时候,需要编译两次.
2. redis换了一个库,之前用的一套感觉很恶心.
3. 其实我的golang水平有限,所以大牛看到有些难以忍受的地方,请赐教.

### 一些别的说明
1.集成mysql,redis.redis 使用 ```github.com/go-redis/redis```,mysql用的是 ```github.com/gohouse/gorose```.
2.有阿里云的oss上传功能.
3.有生成二维码的功能(指定彩色logo).
4.有普罗米修斯的监控.
5.有aes加密解密.
6.有grpc的功能,通信协议使用的protobuf(开关可控).
7.配置文件使用toml格式.
8.用了```gitee.com/johng/gf```的框架的协成池的功能.
### 如有问题,可联系qq:545825965

