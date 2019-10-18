package record

import (
	"encoding/json"
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
func (record *TimerRecord) NotClose() bool {
	return record.Start != "" && record.End == ""
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
	return result
}

func NewOne(timestamp string) {
	rec := TimerRecord{Start: timestamp}
	rec.Date = time.Now().Format("2006-01-02")
	j, _ := json.Marshal(rec)
	database.Set(rec.GetKey(), string(j))
}
