package myredis

import "happylemon/conf"
import (
	goredis "github.com/go-redis/redis"
)

var RedisConn *goredis.Client

func InitRedis() *goredis.Client {
	host := conf.Config.Redis.Host + ":" + conf.Config.Redis.Port
	auth := conf.Config.Redis.Auth
	db := conf.Config.Redis.Db
	RedisConn = goredis.NewClient(&goredis.Options{
		Addr:     host,
		Password: auth, // no password set
		DB:       db,   // use default DB
		PoolSize: 3000,
	})
	return RedisConn
	//pong, err := client.Ping().Result()
	//fmt.Println(pong, err)
	// Output: PONG <nil>
}

func IsRedisNil(err error) bool {
	if err == goredis.Nil {
		return true
	}
	return false
}
