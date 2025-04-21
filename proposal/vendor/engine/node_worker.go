package engine

import (
	"context"
	"google.golang.org/grpc"
	"os/exec"
)

type WorkerNode struct {
	conn	*grpc.ClientConn
	c		NodeServiceClient
}

func ( n *WorkerNode ) Init() (err error) {
	n.conn, err = grpc.Dial("0.0.0.0:30000")
	if err != nil {
		return err
	}

	n.c = NewNodeServerClient(n.conn)

	return nil
}

func ( n *WorkerNode ) Start() {
	_, _ = n.c.ReportStatus(context.Background(), &Request{})

	stream, _ := n.c.AssignTask(context.Background(), &Request{})
	for {

	}
}
