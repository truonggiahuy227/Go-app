package util

import (
	"time"
)

func GetTimestampMinuteCron() time.Time {
	date_now := time.Now()
	sdate_now := date_now.Add(1 * time.Minute).Format("2006-01-02T15:04:00Z07:00")
	layout := "2006-01-02T15:04:05Z07:00"
	now, _ := time.Parse(layout, sdate_now)
	return now
}
