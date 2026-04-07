package utils

import (
  "context"
  rpcxServer "github.com/smallnest/rpcx/server"
  rpcxClient "github.com/smallnest/rpcx/client"
)

// NOTE: repack rpcx framework
// https://github.com/smallnest/rpcx/blob/2ac561ec72aca835b466813ea8353f717ac3162b/client/client.go
// https://doc.rpcx.io/part1/client.html
// https://github.com/rpcxio/rpcx-examples/blob/master/101basic/client/client.go

var gRpcx *rpcxServer.Server
var gRpcxEnable bool
var gRpcxClient *rpcxClient.Client
var gRpcxClientEnable bool

// ------------------------------------------------------------------------------------
// rpcx client
func InitRpcxClient() *rpcxClient.Client {
  gRpcxClient = rpcxClient.NewClient(rpcxClient.DefaultOption)
  gRpcxClientEnable = true
  return gRpcxClient
}

// NOTE: addr ":8080"
func RpcxClientConn(addr string) error {
  if gRpcxClientEnable {
    return gRpcxClient.Connect("tcp", addr)
  }
  return nil
}

func RpcxClientClose()  {
  if gRpcxClientEnable {
    gRpcxClient.Close()
  }
}

// NOTE: args/reply are struct pointer
func RpcxClientCall(name string, method string, args interface{}, reply interface{}) error {
  if gRpcxClientEnable {
    return gRpcxClient.Call(context.Background(), name, method, args, reply)
  }
  return nil
}

// ------------------------------------------------------------------------------------
// rpcx server
func InitRpcxServer() *rpcxServer.Server {
  gRpcx = rpcxServer.NewServer()
  gRpcxEnable = true
  return gRpcx
}

// addr: ":8080"
func RpcxServerStart(addr string) error {
  if gRpcxEnable {
    return gRpcx.Serve("tcp", addr)
  }
  return nil
}

func RpcxClose()  {
  if gRpcxEnable {
    gRpcx.Shutdown(context.Background())
  }
}

// NOTE: obj is service map struct
func RpcxRegister(obj map[string]interface{}) bool {
  if gRpcxEnable && len(obj) > 0 {
    for k, v := range obj {
      gRpcx.RegisterName(k, v, "")
    }
    return true
  }
  return false
}
