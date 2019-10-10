package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	// get system time
	systemTime := time.Now().Format("2006-01-02T15:04:05Z07:00")
	// get net time via http request
	resp, _ := http.Get("http://www.ntsc.ac.cn")
	foo := resp.Header.Get("Date")

	log.Println(systemTime)
	log.Println(foo)
	//save to db //基于文件的数据库

	//search
	//查询某一天最早的时间

	// 历史数据整理
	// 从实时表统计  ，结果 插入历史表
}
