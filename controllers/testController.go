package controllers

import (
	"encoding/json"
	"happylemon/conf"
	"happylemon/lib/encrypt"
	"happylemon/lib/limit"
	"happylemon/mygrpc/apigrpc"
	"happylemon/service"
	"runtime"

	"github.com/gin-gonic/gin"
)

var testSv service.TestService

type TestController struct {
	*BaseController
}

func init() {
	runMode := conf.Config.Server.Mode
	if runMode != "test" {
		return
	}
}

func (self *TestController) Wcd(c *gin.Context) {
	token, _ := c.Get("token")
	needLimit := limit.AccessLimit("wcd", 1, 2)
	if needLimit {
		c.AbortWithStatus(401)
		return
	}
	self.Response(c, conf.CodeOk, map[string]interface{}{
		"token": token,
	})
}

func (self *TestController) Jiami(c *gin.Context) {
	postData, err := c.GetRawData()
	if err != nil {
		c.JSON(200, gin.H{"ret": conf.CodeParaErr})
		return
	}
	ok, okData := encrypt.Encryper(postData)
	if !ok {
		c.JSON(200, gin.H{"ret": conf.CodeAesErr})
		return
	}
	c.JSON(200, gin.H{"ret": conf.CodeOk, "jiami": string(okData)})
}

func (self *TestController) Jiemi(c *gin.Context) {
	postData, err := c.GetRawData()
	if err != nil {
		c.JSON(200, gin.H{"ret": conf.CodeParaErr})
		return
	}
	ok, okData := encrypt.Decryper(postData)
	if !ok {
		c.JSON(200, gin.H{"ret": conf.CodeAesErr})
		return
	}
	c.JSON(200, gin.H{"jiemi": string(okData)})
}

func (self *TestController) PostData(c *gin.Context) {
	retCode, retDat := self.Request(c)
	if retCode != conf.CodeOk {
		self.Response(c, retCode, nil)
	}
	var okData map[string]string
	json.Unmarshal(retDat, &okData)
	self.Response(c, retCode, okData)
}

func (self *TestController) TestDb(c *gin.Context) {
	userinfo := testSv.GetDbInfo()
	self.Response(c, conf.CodeOk, map[string]interface{}{
		"userinfo": userinfo,
	})
}
func (self *TestController) TestRedis(c *gin.Context) {
	userinfo := testSv.GetCacheInfo()
	self.Response(c, conf.CodeOk, map[string]interface{}{
		"userinfo": userinfo,
	})
}
func (self *TestController) Status(c *gin.Context) {
	self.Response(c, conf.CodeOk, map[string]interface{}{
		"NumGoroutine": runtime.NumGoroutine(),
	})
}

func (self *TestController) ApiGrpc(c *gin.Context) {
	para := map[string]interface{}{
		"type": 1,
	}
	paraJson, _ := json.Marshal(para)
	apigrpc.DialGrpc("wcd", string(paraJson))
}

func (self *TestController) Token(c *gin.Context) {
	to := testSv.MakeToken("12233")
	self.Response(c, conf.CodeOk, map[string]interface{}{
		"token": to,
	})
}

func (self *TestController) WcdLog(c *gin.Context) {
	testSv.WcdTestLog("gogogo")
	self.Response(c, conf.CodeOk, nil)
}
