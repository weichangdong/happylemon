package controllers

import (
	"fmt"
	"happylemon/conf"
	"happylemon/entity"
	"happylemon/lib/myredis"
	"happylemon/lib/mytime"
	"os"

	"gitee.com/johng/gf/g/os/grpool"
	"github.com/360EntSecGroup-Skylar/excelize"
	jsoniter "github.com/json-iterator/go"
)

var myjson = jsoniter.ConfigCompatibleWithStandardLibrary

type CmdController struct {
}

func (self *CmdController) CmdSync(act string) {
	switch act {
	case "testLog":
		testSv.WcdTestLog("hello golang")
	default:
		fmt.Println("plz tell me do what!")
		return
	}
	fmt.Println(act + " done")
}

func (self *CmdController) QueueRun() {
	runMode := os.Getenv("RUNMODE")
	cronQueueKey := conf.Config.RedisKey.QueueListKey + runMode
	for {
		grpool.Add(func() {
			executor(cronQueueKey)
		})
		mytime.Usleep(100)
	}
}

func (self *CmdController) ExcelTool() {
	const baseDir = "/Users/weichangdong/Desktop/auto/"
	okDir := baseDir + "user/"
	fileName := "new-user.xlsx"
	xlsx, err := excelize.OpenFile(okDir + fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Get all the rows in the Sheet1.
	rows := xlsx.GetRows("工作表1")
	for k, row := range rows {
		if len(row) != 7 {
			fmt.Println("[error]row num ", "row-num:", k, "row-content:", row)
			continue
		}
	}
}

//真正的执行者
func executor(key string) {
	for {
		rawData, err := myredis.RedisConn.RPop(key).Result()
		if myredis.IsRedisNil(err) {
			break
		}
		inputValue := entity.QueueData{}
		err = myjson.Unmarshal([]byte(rawData), &inputValue)
		if err != nil {
			fmt.Println(err)
			continue
		}
		act := inputValue.Act
		data := inputValue.Data
		switch act {
		case "queue_1":
			fmt.Println("queue_1")
			fmt.Println(data)
			//mytime.Sleep(1)
		}
	}
}
