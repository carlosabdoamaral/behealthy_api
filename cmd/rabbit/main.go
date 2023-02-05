package main

import (
	"github.com/username/.../backend/common"
	"github.com/username/.../backend/internal/grpc"
	"github.com/username/.../backend/internal/rabbit"
)

func main() {
	common.GetEnvVariables()

	grpc.ConnectToGRPCServer()

	rabbit.Connect()
	rabbit.DeclareQueue()
	rabbit.StartConsuming()
}