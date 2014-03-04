package mscrpc

import (
	"net/rpc"
	"net/rpc/jsonrpc"
)

type RpcClient interface {
	Start()
	Call(method string, ret interface{}, v ...interface{}) error 
}

type JsonRpcClient struct {
	client *rpc.Client
}

func NewJsonRpcClient(addr string) (*JsonRpcClient, error) {
	client, err := jsonrpc.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}

	return &JsonRpcClient{client: client}, nil
}

func (c *JsonRpcClient) Call(method string, ret interface{}, v ...interface{}) error {
	args := NewArgs()
	for _, val := range v {
		args.Add(val)
	}

	return c.client.Call(method, args, ret)
}
