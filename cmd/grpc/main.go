package main

import (
	"fmt"

	"github.com/carlosabdoamaral/behealthy_api/common"
	"github.com/carlosabdoamaral/behealthy_api/internal/grpc"
	"github.com/carlosabdoamaral/behealthy_api/internal/persistence"
)

func main() {
	common.GetEnvVariables()

	_, err := persistence.Connect()
	if err != nil {
		fmt.Printf("[!] %s", err.Error())
		return
	}

	grpc.ServeGrpcServer()
}
