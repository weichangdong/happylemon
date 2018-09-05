package middleware

import (
	"fmt"
	"happylemon/conf"
	"happylemon/lib/myredis"
	"happylemon/lib/util"

	"happylemon/lib/prometheus"
	"time"

	"github.com/gin-gonic/gin"
)

func CheckIp() gin.HandlerFunc {
	myfunc := func(c *gin.Context) {
		fromIp := c.ClientIP()
		ipWhiteList := conf.Config.IpWhiteList.Ips
		if !util.IpInArray(fromIp, ipWhiteList) {
			fmt.Println(fromIp)
			c.AbortWithStatus(403)
			return
		}
		c.Next()
	}
	return myfunc
}

func NocheckToken(c *gin.Context) {
	c.Set("httpReqTime", time.Now())
	c.Next()
}
func CheckToken(c *gin.Context) {
	c.Set("httpReqTime", time.Now())

	token := c.Request.Header.Get("APPINFO")
	if token == "" {
		c.AbortWithStatus(401)
		prometheus.HttpCodeCount(c, 401)
		return
	}

	tokenKey := conf.Config.RedisKey.TokenPrefix + token
	tokenUid := conf.Config.RedisKey.TokenUidKey

	uid, _ := myredis.RedisConn.HGet(tokenKey, tokenUid).Result()
	if uid == "" {
		c.AbortWithStatus(401)
		prometheus.HttpCodeCount(c, 401)
		return
	}
	c.Set("uid", uid)
	c.Set("token", token)
	c.Next()
}
