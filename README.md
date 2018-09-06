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
本机mac简单压测了下,压测过程中,没有出现报错(压测过程中,./happylemon -mode=cmd -cmd=queue 一直运行中).
```
~$wrk -c 50 -t 8 -d 30 'http://127.0.0.1:8888/httpstatus?aes=1'
Running 30s test @ http://127.0.0.1:8888/httpstatus?aes=1
  8 threads and 50 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     2.33ms    3.29ms  33.91ms   86.58%
    Req/Sec     4.90k     1.97k   10.43k    73.79%
  1171008 requests in 30.07s, 176.51MB read
Requests/sec:  38945.17
Transfer/sec:      5.87MB
```
redis读取(切记使用的时候不能close,开始close了,压测好多报错,说是Client is close)
```
~$wrk -c 50 -t 8 -d 30 'http://127.0.0.1:8888/redis?aes=1'
Running 30s test @ http://127.0.0.1:8888/redis?aes=1
  8 threads and 50 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     3.40ms    4.69ms  78.50ms   89.05%
    Req/Sec     2.51k   340.33     4.51k    81.42%
  600990 requests in 30.05s, 89.98MB read
Requests/sec:  19999.28
Transfer/sec:      2.99MB
```
mysql读取
```
~$wrk -c 50 -t 8 -d 30 'http://127.0.0.1:8888/db?aes=1'
Running 30s test @ http://127.0.0.1:8888/db?aes=1
  8 threads and 50 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     3.89ms    5.63ms 103.44ms   83.94%
    Req/Sec     1.19k   832.51     3.53k    89.81%
  204011 requests in 30.08s, 114.01MB read
Requests/sec:   6782.53
Transfer/sec:      3.79MB
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
* 使用gin框架构建的一个mvc(controllers,models,entity,service)模式框架,主要是提供api和命令行模式(后台异步执行一些逻辑)执行命令.可通过参数指定运行模式.基于之前开发的一个项目改的,相比之前的项目,做了几点优化.
* 之前命令行模式需要单独go build,然后提供api接口的需要go build一个,等于是两个二进制文件,这样面临的问题有
	+ 命令行执行某个操作完成之后通过channel通知api,但是因为不是一个应用,channel没法通信.
	+ 代码发布的时候,需要编译两次.
* redis换了一个库,之前用的一套感觉很恶心.
* 其实我的golang水平有限,所以大牛看到有些难以忍受的地方,请赐教.

### 一些别的说明
* 集成mysql,redis.redis 使用 ```github.com/go-redis/redis```,mysql用的是 ```github.com/gohouse/gorose```.
* 有阿里云的oss上传功能.
* 有生成二维码的功能(指定彩色logo).
* 有普罗米修斯的监控.
* 有aes加密解密,开关可控不加密.
* 有grpc的功能,通信协议使用的protobuf(开关可控)```protoc --go_out=plugins=grpc:. api.proto```.
* 配置文件使用toml格式.
* 用了```gitee.com/johng/gf```的框架的协成池的功能.
* 使用了concurrentcache,内存级别的缓存,主要应用是接口频率的限制.
### 如有问题,可联系qq:545825965

