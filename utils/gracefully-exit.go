package utils

import (
	"os"
	"os/signal"
	"syscall"
)

// NOTE: gracefully exit
// put AddGracefulExit() func to the end of main()
// utils.AddGracefulExit(func () error {
//   // TODO:
//   return nil
// })  // finish auto exit

type GracefulExitCallback func() error

func AddGracefulExit(fn GracefulExitCallback) { // 编译build后运行 信号捕获就会正常接收了
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM) // 15
	signal.Notify(sig, syscall.SIGINT)  // 2 ctrl+c

	// Block until a signal is received.
	<-sig

	fn() // run self defined function

	os.Exit(0)
}
