package main

import (
	"fmt"

	"github.com/username/.../backend/common"
	"github.com/username/.../backend/internal/grpc"
	"github.com/username/.../backend/internal/persistence"
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