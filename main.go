package main

import (
	"flag"
	"fmt"
	"happylemon/conf"
	"happylemon/controllers"
	"happylemon/lib/log"
	"happylemon/lib/myredis"
	"happylemon/lib/util"
	middle "happylemon/middleware"
	"happylemon/models"
	"happylemon/mygrpc/apigrpc"
	"math"
	"runtime"

	"happylemon/lib/prometheus"

	"github.com/gin-gonic/gin"
)

var runMode, cmd, configFile, method string

func init() {
	flag.StringVar(&runMode, "mode", "cmd", "cmd,http")
	flag.StringVar(&cmd, "cmd", "help", "sync,queue,tool")
	flag.StringVar(&configFile, "conf", "./conf/wcd.toml", "toml config file")
	flag.StringVar(&method, "method", "", "func of tool")
	flag.Parse()
	if !util.FileExists(configFile) {
		panic("config file " + configFile + " no exist!")
	}

	cfg := conf.ReadCfg(configFile)
	log.InitLog(conf.Config.Server.RootPath + conf.Config.Server.LogPath)
	myredis.InitRedis()
	models.InitMysqlPool(cfg)
	if runMode == "http" && conf.Config.Grpc.GrpcSwitch {
		go apigrpc.StartGrpc()
	}

}

func main() {
	if runMode == "cmd" {
		cmdServer(cmd)
	} else {
		httpServer()
	}
}

func cmdServer(cmd string) {
	cmdRouter := new(controllers.CmdController)
	switch cmd {
	case "test":
		fmt.Println("im test")
	case "help":
		helpInfo := `
####################################################
./happylemon -mode=cmd (-conf=conf/wcd.toml) -cmd=sync -method=testLog 命令行模式
./happylemon -mode=http (-conf=conf/wcd.toml) http服务模式
####################################################
		`
		fmt.Println(helpInfo)
	case "sync":
		cmdRouter.CmdSync(method)
	case "queue":
		cmdRouter.QueueRun()
	default:
		fmt.Println(cmd + ":cmd does not exist")
	}
}
func httpServer() {

	NumCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(int(math.Max(float64(NumCPU-1), 1)))

	gin.SetMode(conf.Config.Server.Ginmode)
	router := gin.New()
	router.Use(middle.MyRecoveryWithWriter())
	testRouter := new(controllers.TestController)

	//测试工具
	router.GET("/db", middle.NocheckToken, testRouter.TestDb)
	router.GET("/redis", middle.NocheckToken, testRouter.TestRedis)
	router.POST("/post", middle.NocheckToken, testRouter.PostData)
	router.GET("/testapigrpc", middle.NocheckToken, testRouter.ApiGrpc)
	router.GET("/httpstatus", middle.NocheckToken, testRouter.Status)
	router.GET("/testlog", testRouter.WcdLog)

	//监控统计
	router.GET("/metrics", middle.CheckIp(), prometheus.Handler())
	router.NoRoute(func(c *gin.Context) {
		c.JSON(200, gin.H{"ret": 4})
		prometheus.HttpCodeCount(c, 404)
	})
	fmt.Println("Server Port " + conf.Config.Server.Port)
	router.Run(conf.Config.Server.Port)
}
