package sysload

import (
	"math"
	"runtime"
	"time"
)

type LoadAvgInfo struct {
	One     float64
	Five    float64
	Fifteen float64
	Time    time.Time
}

type MemInfo struct {
	MemCached  uint64
	MemTotal   uint64
	MemUsed    uint64
	MemFree    uint64
	SwapCached uint64
	SwapTotal  uint64
	SwapUsed   uint64
	SwapFree   uint64
	Time       time.Time
}

func GetLoadAvgInfo() (load LoadAvgInfo) {
	return getLoadAvgInfo()
}

func GetMemInfo() (mem MemInfo) {
	return getMemInfo()
}

// 简单的系统负载计算方法，仅考虑内存和CPU负载情况
// CPU负载按照1/5/15分钟分别占60%/30%/10%
// 在整个负载考量数据中，CPU占80%，内存为20%
func GetSysLoad() int {
	loadAvg := GetLoadAvgInfo()
	memInfo := GetMemInfo()
	memLoad := float64(memInfo.MemUsed) / float64(memInfo.MemTotal) // 暂时不考虑Swap
	cpuLoad := (loadAvg.One*0.6 + loadAvg.Five*0.3 + loadAvg.Fifteen*0.1) / float64(runtime.NumCPU())
	load := (cpuLoad*0.8 + memLoad*0.2) * 100
	return int(math.Ceil(load))
}
