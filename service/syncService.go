package service

import (
	"encoding/json"
	"happylemon/conf"
	"happylemon/entity"
	"happylemon/lib/myredis"
	"os"
)

type SyncService struct{}

//定时执行(crontab)的队列
func (self *SyncService) AddToCron(queueData map[string]interface{}) {
	runMode := os.Getenv("RUNMODE")
	cronQueueKey := conf.Config.RedisKey.CronListKey + runMode
	okData, _ := json.Marshal(queueData)
	myredis.RedisConn.LPush(cronQueueKey, string(okData))
}

//几乎等于同步的执行队列
func (self *SyncService) AddToQueue(queueData entity.QueueData) {
	runMode := os.Getenv("RUNMODE")
	cronQueueKey := conf.Config.RedisKey.QueueListKey + runMode
	okData, _ := json.Marshal(queueData)
	myredis.RedisConn.LPush(cronQueueKey, string(okData))
}
