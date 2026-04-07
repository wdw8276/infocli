package utils

import (
	"github.com/wangdiwen/gopool"
)

// default queue size 1000000
// pollNum is maxWorkers == minWorkers
func InitPool(pollNum int) gopool.GoPool {
	return gopool.NewGoPool(pollNum)
}

// example
// https://pkg.go.dev/github.com/devchat-ai/gopool#section-readme

// gNum := 0
// logger.Println("gNum=", gNum)

// go func() {
// 	for true {
// 		queueSize := p.GetTaskQueueSize()
// 		logger.Println("task queue size: ", queueSize)
// 		workerCount := p.GetWorkerCount()
// 		logger.Println("task worker count: ", workerCount)
// 		runCount := p.Running()
// 		logger.Println("task running count: ", runCount)
// 		utils.SleepMilli(250)
// 	}
// }()

// p := utils.InitPool(10)
// defer p.Release()
// p.AddTask(func() (interface{}, error) {
// 	logger.Println("task 1")
// 	gNum += 1
// 	return nil, nil
// })
// p.Wait()

// logger.Println("gNum=", gNum)
