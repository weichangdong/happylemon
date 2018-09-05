package limit

import (
	"time"

	"encoding/hex"

	"github.com/jiachuhuang/concurrentcache"
)

var cache, _ = concurrentcache.NewConcurrentCache(256, 10240)

func AccessLimit(key string, limitNum int, limitTime int) bool {
	oldNum, _ := cache.Get(key)
	var oldOkNum = 0
	if oldNum != nil {
		oldOkNum = oldNum.(int)
	}
	if oldOkNum >= limitNum {
		//c.AbortWithStatus(401)
		//c.JSON(401, gin.H{})
		return true
	} else {
		newNum := oldOkNum + 1
		cache.Set(key, newNum, time.Duration(limitTime)*time.Second)
	}
	return false
}

func AeskeyCache(aeskey string) (key []byte) {
	key_tmp, _ := cache.Get("aeskey")
	if key_tmp == nil {
		key, _ = hex.DecodeString(aeskey)
		cache.Set("aeskey", key, 7*24*time.Hour)
	} else {
		key = key_tmp.([]byte)
	}
	return
}

func SetDataMemory(key string, value interface{}, expireTime time.Duration) {
	cache.Set(key, value, expireTime*time.Second)
}

func GetDataMemory(key string) (bool, interface{}) {
	value, err := cache.Get(key)
	if err != nil || value == nil {
		return false, nil
	}
	return true, value
}
