package grpc

import (
	"fmt"
	"net"

	"github.com/carlosabdoamaral/behealthy_api/common"
	"google.golang.org/grpc"
)

type TemplateServer struct {
	// pb.UnimplementedAccountServiceServer

}

func ServeGrpcServer() {
	fmt.Println("[*] Starting GRPC")
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		fmt.Printf("[!] %s", err.Error())
		return
	}

	common.GrpcServer = grpc.NewServer()
	// pb.RegisterAccountServiceServer(common.GrpcServer, &TemplateServer{})

	fmt.Printf("[*] Server listening on %v", lis.Addr())
	err = common.GrpcServer.Serve(lis)
	if err != nil {
		fmt.Printf("[!] Failed to serve, error: %s", err.Error())
		return
	}
}
