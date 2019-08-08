package kbutils

import "time"

var serverRunTime int64

func init() {
	serverRunTime = time.Now().Unix()
}

func GetRunSecond() int64 {
	return time.Now().Unix() - serverRunTime
}
