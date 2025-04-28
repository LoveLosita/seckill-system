package utils

import "time"

func StrTimeToUnix(startTime, endTime string) (int64, int64, error) {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	startTimeUnix, err := time.ParseInLocation("2006-01-02 15:04:05", startTime, loc)
	if err != nil {
		return 0, 0, err
	}
	endTimeUnix, err := time.ParseInLocation("2006-01-02 15:04:05", endTime, loc)
	if err != nil {
		return 0, 0, err
	}
	return startTimeUnix.Unix(), endTimeUnix.Unix(), nil
}
