#!/usr/bin/env bash

if [[ -z ~/go/bin/proto-gen-go-grpc ]]; then
    echo Please make sure Go is installed, and run
    echo -e '\t$ go install google.golang.org/protobuf/cmd/protoc-gen-go@latest'
    echo -e '\t$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest'
fi

# add protoc to path
if [[ -z ~/go/bin/proto-gen-go ]]; then
    ln -s $HOME/go/bin/protoc-gen-go-grpc $HOME/go/bin/protoc-gen-go
fi

# include gopath if not already included
export GOPATH=$HOME/go
case :$PATH:
    in  *:$GOPATH/bin:*) ;;
    *) export PATH=$PATH:$GOPATH/bin;;
esac

protoc --go-grpc_out=./vendor/engine \
    node.proto
