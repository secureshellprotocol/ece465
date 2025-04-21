package engine

import (
	"context"
)

type NodeGrpcServer struct {
	command_channel chan string
}

func ( n NodeGrpcServer ) ReportStatus(ctx context.Context, rqst *Request) (*Response, error) {
	return &Response{Data: "ok"}, nil
}

func ( n NodeGrpcServer ) AssignTask(rqst *Request, server *NodeService_AssignTaskServer) error {
	for {
		select {
		case cmd := <-n.command_channel:
			// send command to worker node
			if err := server.Send(&Response{Data: cmd}); err != nil {
				return err
			}
		}
	}
}

var server *NodeGrpcServer

func GetNodeGrpcServer() *NodeGrpcServer {
	if server == nil {
		server = &NodeGrpcServer {
			command_channel: make(chan string),
		}
	}
	return server
}
