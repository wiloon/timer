package record

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"github.com/wiloon/pingd-utils/utils"
	"time"
	"timer/database"
)

const timerRecordSuffix = "_record"

type TimerRecord struct {
	Date  string
	Start string
	End   string
}

func (record *TimerRecord) IsEmpty() bool {
	return record.Start == "" && record.End == ""
}
func (record *TimerRecord) IsClosed() bool {
	closed := record.Start != "" && record.End != ""
	log.Infof("record is closed: %v", closed)
	return closed
}
func (record *TimerRecord) GetKey() string {
	return record.Date + timerRecordSuffix
}
func (record *TimerRecord) Close(s string) {
	record.End = s
	j, _ := json.Marshal(record)
	database.Set(record.GetKey(), string(j))
}
func GetRecordByDate(date string) TimerRecord {
	var result TimerRecord
	value := database.Get(date + timerRecordSuffix)
	if value != "" {
		_ = json.Unmarshal([]byte(value), &result)
	}
	log.Infof("get record by date, date: %v, result: %v", date, result)
	return result
}

func NewOne(timestamp time.Time) {
	rec := TimerRecord{Start: utils.DateToStringYMDHMSZ(timestamp)}
	rec.Date = utils.DateToStringYMD(timestamp)
	j, _ := json.Marshal(rec)
	database.Set(rec.GetKey(), string(j))
}
