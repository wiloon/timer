package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/wiloon/pingd-log/logconfig"
	"github.com/wiloon/pingd-utils/utils"
	"net/http"
	"time"
	"timer/database"
	"timer/heartbeat"
	"timer/record"
)

func main() {
	logconfig.Init()
	// get system time
	systemTime := time.Now()
	systemTimeStr := systemTime.Format("2006-01-02T15:04:05Z07:00")
	systemTimeStrDay := systemTime.Format("2006-01-02")
	// get net time via http request
	resp, _ := http.Get("http://www.ntsc.ac.cn")
	netTimeOriginal := resp.Header.Get("Date")
	netTime := utils.StringToDateRFC1123(netTimeOriginal)
	loc, _ := time.LoadLocation("Local")
	netTime = netTime.In(loc)
	netTimeStr := netTime.Format("2006-01-02T15:04:05Z07:00")
	log.Infof("system time: %v, net time: %v", systemTimeStr, netTimeStr)

	// 先查一下昨天的记录有没有
	d, _ := time.ParseDuration("-24h")
	yesterday := systemTime.Add(d)
	yesterdayStr := yesterday.Format("2006-01-02")

	yesterdayRecord := record.GetRecordByDate(yesterdayStr)

	if yesterdayRecord.NotClose() {
		// 如果 昨天的记录没关掉，查询昨天最后一次心跳
		yesterdayHeartbeat := heartbeat.GetByDate(yesterdayStr)
		if yesterdayHeartbeat != "" {
			//填充昨天的记录
			yesterdayRecord.Close(yesterdayHeartbeat)
		}
	}

	//检查今天的心跳是否存在
	todayHeartbeat := heartbeat.GetByDate(systemTimeStrDay)
	if todayHeartbeat == "" {
		//今天的第一次心跳作为记录的开始时间
		record.NewOne(systemTimeStr)

	}

	//log.Infof("yesterday end time, key: %v, value: %v", yesterdayStr, yesterdayEndTime)
	//if yesterdayEndTime != "" {
	//
	//}

	database.Set(systemTimeStrDay, systemTimeStr)

	value := database.Get(systemTimeStrDay)
	log.Info("value:" + value)

	//search
	//查询某一天最早的时间

	// 历史数据整理
	// 从实时表统计  ，结果 插入历史表
	log.Info("end")
}
