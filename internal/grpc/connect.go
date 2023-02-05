package grpc

import (
	"flag"
	"log"

	"github.com/carlosabdoamaral/behealthy_api/common"
	pb "github.com/carlosabdoamaral/behealthy_api/protodefs/gen/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ConnectToGRPCServer() *grpc.ClientConn {
	var Addr = flag.String("addr", "localhost:50051", "the address to connect to")

	flag.Parse()

	conn, err := grpc.Dial(*Addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Error connecting to %v: %v", Addr, err)
	}

	common.GrpcConn = conn
	common.AccountServiceClient = pb.NewAccountServiceClient(conn)
	common.AuthServiceClient = pb.NewAuthServiceClient(conn)
	common.ExerciseServiceClient = pb.NewExerciseServiceClient(conn)
	common.ReportServiceClient = pb.NewReportServiceClient(conn)
	common.WorkoutServiceClient = pb.NewWorkoutServiceClient(conn)

	return conn
}
