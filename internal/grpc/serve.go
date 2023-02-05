package grpc

import (
	"fmt"
	"net"

	"github.com/carlosabdoamaral/behealthy_api/common"
	pb "github.com/carlosabdoamaral/behealthy_api/protodefs/gen/proto"
	"google.golang.org/grpc"
)

type AccountServer struct {
	pb.UnimplementedAccountServiceServer
}

type AuthServer struct {
	pb.UnimplementedAuthServiceServer
}

type ExerciseServer struct {
	pb.UnimplementedExerciseServiceServer
}

type ReportServer struct {
	pb.UnimplementedReportServiceServer
}

type WorkoutServer struct {
	pb.UnimplementedWorkoutServiceServer
}

func ServeGrpcServer() {
	fmt.Println("[*] Starting GRPC")
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		fmt.Printf("[!] %s", err.Error())
		return
	}

	common.GrpcServer = grpc.NewServer()
	pb.RegisterAccountServiceServer(common.GrpcServer, &AccountServer{})
	pb.RegisterAuthServiceServer(common.GrpcServer, &AuthServer{})
	pb.RegisterExerciseServiceServer(common.GrpcServer, &ExerciseServer{})
	pb.RegisterReportServiceServer(common.GrpcServer, &ReportServer{})
	pb.RegisterWorkoutServiceServer(common.GrpcServer, &WorkoutServer{})

	fmt.Printf("[*] Server listening on %v", lis.Addr())
	err = common.GrpcServer.Serve(lis)
	if err != nil {
		fmt.Printf("[!] Failed to serve, error: %s", err.Error())
		return
	}
}
