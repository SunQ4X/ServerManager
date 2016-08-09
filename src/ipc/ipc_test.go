package ipc

import (
	//"fmt"
	"testing"
)

type EchoServer struct {
	name string "name"
}

func (server *EchoServer) Name() string {
	return server.name
}

func (server *EchoServer) Handle(method, params string) *Response {
	return &Response{"200", "Echo:" + params}
}

func TestIpc(t *testing.T) {
	server1 := NewIpcServer(&EchoServer{"server1"})
	server2 := NewIpcServer(&EchoServer{"server2"})

	client1 := NewIpcClient(server1)
	client2 := NewIpcClient(server2)

	resp1, _ := client1.Call("call", "From client1")
	resp2, _ := client2.Call("call", "From client2")

	if resp1.Body != "Echo:From client1" || resp2.Body != "Echo:From client2" {
		t.Error("resp error. resp1:", resp1, "resp2:", resp2)
	}
}
