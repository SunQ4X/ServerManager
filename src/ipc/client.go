package ipc

import (
	"encoding/json"
	//"time"
)

type IpcClient struct {
	conn chan string
}

func NewIpcClient(server *IpcServer) *IpcClient {
	conn := server.Connect()
	return &IpcClient{conn}
}

func (client *IpcClient) Call(method, params string) (*Response, error) {
	req := &Request{method, params}

	request, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	client.conn <- string(request)
	response := <-client.conn

	var resp Response
	err = json.Unmarshal([]byte(response), &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
