package engine

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"net"
	"net/http"
)

type ControllerNode struct {
	api			*gin.Engine
	ln			net.Listener
	svr			*grpc.Server
	node_svr	NodeGrpcServer
}

func ( n *ControllerNode ) Init (err error) {
	n.ln, err = net.Listen("tcp", ":30000")
	if err != nil {
		return err
	}

	n.svr = grpc.NewServer()

	n.node_svr = GetNodeGrpcServer()

	RegisterNodeServicesServer(n.svr, n.node_svr)

	// frontend api -- TODO break into own functions
	n.api = gin.Default()
	n.api.POST("/tasks", func(c *gin.Context) {
		var payload struct {
			cmd string `json:"cmd"`
		}
		if err := c.ShouldBindJSON(&payload); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return err
		}

		n.node_svr.command_channel <- payload.cmd
		c.AbortWithStatus(http.StatusTeapot) // ????
	})

	return nil
}

func ( n *ControllerNode ) Start() {
	// grpc
	go n.svr.Serve(n.ln)


	// api
	_ = n.api.run(":8080")

	n.svr.Stop()
}

var node *ControllerNode

func GetControllerNode() *ControllerNode {
	if node == nil {
		node = &ControllerNode{}

		if err := node.Init(); err != nil { 
			panic(err)
		}
	}

	return node
}
