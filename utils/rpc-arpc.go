package utils

import (
	"context"
	"net"
	"sync"
	"time"

	arpcTool "github.com/lesismal/arpc"
	arpcRouter "github.com/lesismal/arpc/extension/middleware/router"
	arpcLogger "github.com/lesismal/arpc/log"
)

// NOTE: repack arpc, support call, notify, broadcast

type ArpcContext = arpcTool.Context // rename Context
type ArpcFunc = func(*ArpcContext)  // rename func

// server pointer
var gArpc *arpcTool.Server
var gArpcMux *sync.RWMutex
var gArpcGraceful *arpcRouter.Graceful
var gArpcClientMap = make(map[*arpcTool.Client]struct{})
var gArpcEnable bool
var gArpcRunning bool

// client pointer
var gArpcClient *arpcTool.Client
var gArpcClientEnable bool

// set logger level, iota从0开始递增
// 0 LevelAll = iota
// // LevelDebug logs are usually disabled in production.
// 1 LevelDebug
// // LevelInfo is the default logging priority.
// 2 LevelInfo
// // LevelWarn .
// 3 LevelWarn
// // LevelError .
// 4 LevelError
// // LevelNone disables all logs.
// 5 LevelNone
func ArpcSetLoggerLevel(level int) {
	arpcLogger.SetLevel(level)
}

func ArpcSetLoggerFormat() {
	arpcLogger.TimeFormat = "2006/01/02 15:04:05" // 这里修改arpc自带的日志格式 显示到秒即可
}

// ------------------------------------------------------------------------------------
// client
func InitArpcClient(addr string) error {
	c, err := arpcTool.NewClient(func() (net.Conn, error) {
		return net.DialTimeout("tcp", addr, time.Second*3)
	})
	if err != nil {
		gArpcClientEnable = false
	} else {
		gArpcClientEnable = true
		gArpcClient = c
	}
	return err
}

func ArpcClientClose() {
	if gArpcClientEnable {
		gArpcClient.Stop()
	}
}

func ArpcClientRegister(method string, fn ArpcFunc) bool {
	if gArpcClientEnable {
		gArpcClient.Handler.Handle(method, fn)
		return true
	}
	return false
}

// NOTE: v is pointer of *string or *[]byte
func ArpcClientBind(ctx *ArpcContext, v interface{}) error {
	if gArpcClientEnable {
		return ctx.Bind(v)
	}
	return nil
}

// NOTE: req, resp are pointer *[]byte or *string
func ArpcClientCall(method string, req interface{}, resp interface{}) error {
	if gArpcClientEnable {
		return gArpcClient.Call(method, req, resp, time.Second*2)
	}
	return nil
}

// ------------------------------------------------------------------------------------
// server
func InitArpcServer() *arpcTool.Server {
	gArpcMux = new(sync.RWMutex)

	gArpc = arpcTool.NewServer()
	gArpcGraceful = &arpcRouter.Graceful{}
	gArpc.Handler.Use(gArpcGraceful.Handler())

	gArpc.Handler.HandleDisconnected(func(c *arpcTool.Client) {
		gArpcMux.Lock()
		delete(gArpcClientMap, c)
		gArpcMux.Unlock()
	})
	gArpc.Handler.HandleConnected(func(c *arpcTool.Client) {
		gArpcMux.Lock()
		gArpcClientMap[c] = struct{}{}
		gArpcMux.Unlock()
	})

	gArpcEnable = true
	gArpcRunning = false
	return gArpc
}

// addr: ":8080"
func ArpcServerStart(addr string) error {
	if gArpcEnable {
		gArpcRunning = true
		err := gArpc.Run(addr)
		if err != nil {
			gArpcRunning = false
		}
	}
	return nil
}

func ArpcClose() {
	if gArpcEnable && gArpcRunning {
		gArpcGraceful.Shutdown()
		gArpc.Shutdown(context.Background())
	}
}

func ArpcServerStatus() bool {
	if gArpcEnable && gArpcRunning {
		return true
	}
	return false
}

func ArpcRegister(method string, fn ArpcFunc) bool {
	if gArpcEnable {
		gArpc.Handler.Handle(method, fn)
		return true
	}
	return false
}

// NOTE: async send to client
func ArpcSetAsync(v bool) {
	if gArpcEnable {
		gArpc.Handler.SetAsyncWrite(v)
	}
}

// NOTE: v is pointer, support *[]byte or *string
func ArpcSend(ctx *ArpcContext, v interface{}) error {
	if gArpcEnable {
		return ctx.Write(v) // block
	}
	return nil
}

// NOTE: v support *[]byte or *string
func ArpcBind(ctx *ArpcContext, v interface{}) error {
	if gArpcEnable {
		return ctx.Bind(v)
	}
	return nil
}

// NOTE: v support *[]byte or *string
func ArpcNotifyNonblock(ctx *ArpcContext, method string, v interface{}) error {
	return ctx.Client.Notify(method, v, arpcTool.TimeZero)
}
func ArpcNotifyBlock(ctx *ArpcContext, method string, v interface{}) error {
	return ctx.Client.Notify(method, v, arpcTool.TimeForever)
}

// NOTE: v support *[]byte or *string
func ArpcSendBroadcast(method string, v interface{}) bool {
	if gArpcEnable {
		gArpc.Broadcast(method, v)
		return true
	}
	return false
}

func ArpcAddClient(c *ArpcContext) bool {
	if gArpcEnable {
		gArpcMux.Lock()
		gArpcClientMap[c.Client] = struct{}{}
		gArpcMux.Unlock()
	}
	return false
}

func ArpcClientCount() int {
	return len(gArpcClientMap)
}
