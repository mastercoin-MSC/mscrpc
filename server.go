package mscrpc 

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type RpcServer interface {
	Start()
	Stop()
}

type JsonRpcServer struct {
	quit chan bool
	listener net.Listener
}

func NewJsonRpcServer() *JsonRpcServer {
	l, err := net.Listen("tcp", ":12345")
	if err != nil {
		log.Fatal("Error starting JSON-RPC")
	}

	return &JsonRpcServer{
		listener: l,
		quit: make(chan bool),
	}
}

func (s *JsonRpcServer) startHandler() {
out:
	for {
		select {
		case <- s.quit:
			s.listener.Close()
			break out
		}
	}

	log.Println("[JSON] Shutdown JSON-RPC server")
}

func (s *JsonRpcServer) Start() {
	log.Println("[JSON] Starting JSON-RPC server")
	RegisterPackagesRpcPackages()
	rpc.HandleHTTP()

	// Start serving
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go jsonrpc.ServeConn(conn)
	}
}

func (s *JsonRpcServer) Stop() {
	close(s.quit)
}
