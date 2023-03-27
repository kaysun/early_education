package util

import "time"

// TimeToStr 将时间转成字符串
func TimeToStr(theTime time.Time) string {
	return theTime.Format("2006-01-02 15:04:05")
}

// StrToTime 字符串转时间
func StrToTime(timeStr string) time.Time {
	loc, _ := time.LoadLocation("Local")
	theTime, err := time.ParseInLocation("2006-01-02 15:04:05", timeStr, loc)
	if err != nil {
		return time.Time{}
	}
	return theTime
}
