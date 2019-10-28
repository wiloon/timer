package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/wiloon/pingd-log/logconfig"
	"github.com/wiloon/pingd-utils/utils"
	"net/http"
	"time"
	"timer/heartbeat"
	"timer/record"
)

var cachedDateYMD string

func main() {
	logconfig.Init()

	ticker := time.NewTicker(10 * time.Second)
	for ; true; <-ticker.C {
		// get system time
		systemTime := time.Now()

		// get net time via http request
		resp, _ := http.Get("http://www.ntsc.ac.cn")
		netTimeOriginal := resp.Header.Get("Date")
		netTime := utils.StringToDateRFC1123(netTimeOriginal)
		loc, _ := time.LoadLocation("Local")
		netTime = netTime.In(loc)
		netTimeStr := netTime.Format("2006-01-02T15:04:05Z07:00")
		log.Infof("system time: %v, net time: %v", systemTime, netTimeStr)
		timerTask(systemTime)
	}

}

func timerTask(currentTime time.Time) {
	log.Infof("timer task start: %v", currentTime)

	currentTimeYMD := currentTime.Format("2006-01-02")
	currentTimeYMDHMS := currentTime.Format("2006-01-02T15:04:05Z07:00")

	if cachedDateYMD != currentTimeYMD {
		cachedDateYMD = currentTimeYMD

		// 先查一下昨天的记录有没有
		d, _ := time.ParseDuration("-24h")
		yesterday := currentTime.Add(d)
		yesterdayStrYMD := yesterday.Format("2006-01-02")

		yesterdayRecord := record.GetRecordByDate(yesterdayStrYMD)

		if !yesterdayRecord.IsClosed() {
			// 如果 昨天的记录没关掉，查询昨天最后一次心跳
			yesterdayHeartbeat := heartbeat.GetByDate(yesterdayStrYMD)
			if yesterdayHeartbeat != "" {
				//填充昨天的记录
				yesterdayRecord.Close(yesterdayHeartbeat)
				log.Infof("close yesterday record, end time: %v", yesterdayHeartbeat)
			}
		}
	}

	//检查今天的心跳是否存在
	todayHeartbeat := heartbeat.GetByDate(currentTimeYMD)
	if todayHeartbeat == "" {
		//今天的第一次心跳作为记录的开始时间
		record.NewOne(currentTime)
		log.Infof("today heartbeat not exist, create new record, start time: %v", currentTimeYMDHMS)
	}
	heartbeat.Update(currentTimeYMD, currentTimeYMDHMS)
	log.Infof("update heartbeat: %v", currentTimeYMDHMS)
	log.Info("timer task end")
}
