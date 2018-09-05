// +build !linux

package sysload

import (
	"happylemon/lib/log"
	"runtime"
	"time"
)

func getLoadAvgInfo() (load LoadAvgInfo) {
	load.Time = time.Now()
	log.InfoLog("LoadAvg is not implemented in OS: " + runtime.GOOS + "\n")
	return
}

func getMemInfo() (load MemInfo) {
	load.Time = time.Now()
	log.InfoLog("meminfo is not implemented in OS:" + runtime.GOOS + "\n")
	return
}
