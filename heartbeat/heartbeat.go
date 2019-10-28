package heartbeat

import (
	log "github.com/sirupsen/logrus"
	"timer/database"
)

const heartBeatKeySuffix = "_heartbeat"

func GetByDate(date string) string {
	heartBeat := database.Get(date + "_heartbeat")
	log.Infof("get heartbeat by date, date: %v, result: %v", date, heartBeat)
	return heartBeat
}

func Update(dayStr, timestamp string) {
	database.Set(dayStr+"_heartbeat", timestamp)
}

func getDateStr() {
	return
}
