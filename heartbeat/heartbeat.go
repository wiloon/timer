package heartbeat

import "timer/database"

const heartBeatKeySuffix = "_heartbeat"

func GetByDate(date string) string {
	return database.Get(date + "_heartbeat")
}
