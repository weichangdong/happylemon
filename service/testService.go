package service

import (
	"happylemon/lib/log"
	"happylemon/lib/myredis"
	"happylemon/lib/token"
	"happylemon/models"
)

type TestService struct{}

var testMd models.TestModel

func (self *TestService) GetDbInfo() map[string]interface{} {
	userinfo := testMd.ReadInfo()
	return userinfo
}

func (self *TestService) GetCacheInfo() string {
	red := myredis.RedisConn
	val2, _ := red.Get("wcd").Result()
	return val2
}
func (self *TestService) DealStatus(act, para string) string {
	return "act:" + act + " para:" + para
}
func (self *TestService) MakeToken(uid string) string {
	return token.GenNewToken()
}

func (self *TestService) WcdTestLog(str string) {
	log.ErrorLog(str)
	log.InfoLog(str)
}
