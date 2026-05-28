package utils

import (
  "time"
)

func GetNowTime() string {
  return time.Now().Format("2006-01-02 03:04:05")  // format timestamp by golang birth
}

func SleepSec(interval int) {
	time.Sleep(time.Duration(interval) * time.Second)
}

func SleepMilli(interval int) {
	time.Sleep(time.Duration(interval) * time.Millisecond)
}
