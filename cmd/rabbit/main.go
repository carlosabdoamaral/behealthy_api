package main

import (
	"github.com/carlosabdoamaral/behealthy_api/common"
	"github.com/carlosabdoamaral/behealthy_api/internal/grpc"
	"github.com/carlosabdoamaral/behealthy_api/internal/rabbit"
)

func main() {
	common.GetEnvVariables()

	grpc.ConnectToGRPCServer()

	rabbit.Connect()
	rabbit.DeclareQueue()
	rabbit.StartConsuming()
}
