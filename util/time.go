package util

import "time"

const layout = "2006-01-02T15:04:05"

func StringToDate(value string) time.Time {
	convertedDate, _ := time.Parse(layout, value)
	return convertedDate
}
