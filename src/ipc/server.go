package ipc

import (
	"encoding/json"
	"fmt"
)

type Request struct {
	Method string "method"
	Params string "params"
}

type Response struct {
	Code string "code"
	Body string "body"
}

type Server interface {
	Name() string
	Handle(method, params string) *Response
}

type IpcServer struct {
	Server
}

func NewIpcServer(server Server) *IpcServer {
	return &IpcServer{server}
}

func (server *IpcServer) Connect() chan string {
	session := make(chan string)
	go func(c chan string) {
		for {
			request := <-c

			if request == "CLOSE" {
				break
			}

			var req Request
			err := json.Unmarshal([]byte(request), &req)
			if err != nil {
				fmt.Println("Invalid request format:", request)
				continue
			}

			resp := server.Handle(req.Method, req.Params)

			response, err := json.Marshal(resp)
			if err != nil {
				fmt.Println("Response marshal failed.")
			}
			c <- string(response)
		}

		fmt.Println("Session closed.")
	}(session)

	fmt.Println("A new session has been created successfully.")
	return session
}
