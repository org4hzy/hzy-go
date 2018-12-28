package util

import "time"

const dayInter int64 = 86400

var timezone int64 = 8

// Launcher timer default setting
func Launcher() {
	time.LoadLocation("Asia/Shanghai")
}

// GetTs unix timestamp ms
func GetTs() int64 {
	return time.Now().UnixNano() / 1e6
}

// GetZeroTs 0点时间戳 s
func GetZeroTs() int64 {
	now := time.Now().Unix()
	return now - (now+timezone*3600)%dayInter
}

// IsToday 今天的时间戳 ts is ms
func IsToday(ts int64) bool {
	ts /= 1000
	return GetZeroTs() == (ts - (ts+timezone*3600)%dayInter)
}


