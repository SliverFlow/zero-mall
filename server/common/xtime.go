package common

import "time"

const TimeFormat = "2006-01-02 15:04:05"

// FormatTime 将字符串转换为时间
func FormatTime(timeStr string) (time.Time, error) {
	return time.Parse(TimeFormat, timeStr)
}

// ParseTime 将时间转换为字符串
func ParseTime(t time.Time) string {
	return t.Format(TimeFormat)
}
