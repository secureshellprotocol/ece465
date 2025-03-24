package core

import (
	"context"
)

type GrpcServer struct {
	command_channel chan string
}

func ( n GrpcServer ) ReportStatus(ctx context.Context, rqst *Request) (*Response, error) {
	return &Response{Data: "ok"}, nil
}

func ( n GrpcServer ) AssignTask(rqst *Request, server *NodeService_AssignTaskServer) error {
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

var server *GrpcServer

func GetGrpcServer() *GrpcServer {
	if server == nil {
		server = &GrpcServer {
			command_channel: make(chan string),
		}
	}
	return server
}
