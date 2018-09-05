// +build linux

package sysload

import (
	"bufio"
	"bytes"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

func getLoadAvgInfo() (load LoadAvgInfo) {
	procf := "/proc/loadavg" // Linux proc file system
	content, err := ioutil.ReadFile(procf)
	load.Time = time.Now()
	if err != nil {
		return
	}
	reader := bufio.NewReader(bytes.NewBuffer(content))
	line, _, _ := reader.ReadLine()
	fields := strings.Fields(string(line))
	if len(fields) < 3 {
		return
	}
	load.One, _ = strconv.ParseFloat(fields[0], 64)
	load.Five, _ = strconv.ParseFloat(fields[1], 64)
	load.Fifteen, _ = strconv.ParseFloat(fields[2], 64)
	return
}

func getMemInfo() (mem MemInfo) {
	procf := "/proc/meminfo" // Linux proc file
	content, err := ioutil.ReadFile(procf)
	mem.Time = time.Now()
	if err != nil {
		return
	}
	reader := bufio.NewReader(bytes.NewBuffer(content))
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if idx := bytes.IndexByte(line, ':'); idx > 0 {
			line[idx] = ' '
		}
		fields := strings.Fields(string(line))
		if len(fields) != 3 {
			break
		}
		val := strconv.ParseUint(fields[1], 10, 64)
		switch fields[0] {
		case "MemCached":
			mem.MemCached = val
		case "MemTotal":
			mem.MemTotal = val
		case "MemFree":
			mem.MemFree = val
		case "SwapCached":
			mem.SwapCached = val
		case "SwapTotal":
			mem.SwapTotal = val
		case "SwapFree":
			mem.SwapFree = val
		}
	}
	mem.MemUsed = mem.MemTotal - mem.MemFree
	mem.SwapUsed = mem.SwapTotal - mem.SwapFree
	return
}
