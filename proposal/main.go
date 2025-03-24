package main

import (
    "core"
    "os"
)

func main() {
    nodeType := os.Args[0]
    switch nodeType {
    case "master":
        core.GetMasterNode().Start()
    case "worker":
        core.GetWorkerNode().Start()
    default:
        panic("invalid node type")
    }
}
