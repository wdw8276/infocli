package utils

import (
  "time"
)

func GetNowTime() string {
  return time.Now().Format("2006-01-02 03:04:05")  // format timestamp by golang birth
}

// 睡眠 秒
func SleepSec(interval int)  {
  time.Sleep(time.Duration(interval) * time.Second)
}

// 睡眠 毫秒
func SleepMilli(interval int)  {
  time.Sleep(time.Duration(interval) * time.Millisecond)
}
