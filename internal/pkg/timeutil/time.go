package timeutil

import "time"

const (
	DateFormat    = "2006-01-02"
	DateDayFormat = "20060102"
	TimeFormat    = "15:04:05"
)

func DateTime() string {
	return time.Now().Format(time.DateTime)
}

func Date() string {
	return time.Now().Format(DateFormat)
}

func Location() *time.Location {
	lo, _ := time.LoadLocation("Asia/Shanghai")
	return lo
}

func DayStartDateTime() string {
	return Date() + " 00:00:00"
}

func DayEndDateTime() string {
	return Date() + " 23:59:59"
}

func DayStartTime() int64 {
	t, _ := time.ParseInLocation(time.DateTime, DayStartDateTime(), Location())

	return t.Unix()
}

func DayEndTime() int64 {
	t, _ := time.ParseInLocation(time.DateTime, DayEndDateTime(), Location())
	return t.Unix()
}

// ParseDateTime 解析 Datetime 格式字符串为 Time
func ParseDateTime(datetime string) time.Time {
	t, _ := time.ParseInLocation(time.DateTime, datetime, Location())
	return t
}

func FormatDatetime(t time.Time) string {
	return t.Format(time.DateTime)
}

func IsDateTime(datetime string) bool {
	_, err := time.Parse(time.DateTime, datetime)
	return err == nil
}

func IsDate(date string) bool {
	_, err := time.Parse(DateFormat, date)
	return err == nil
}
