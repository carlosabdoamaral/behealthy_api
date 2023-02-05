package common

import (
	"database/sql"

	pb "github.com/carlosabdoamaral/behealthy_api/protodefs/gen/proto"
	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
	"google.golang.org/grpc"
)

var (
	MockModeActivated bool = false
)

var (
	ApiPort string      = "8080"
	Router  *gin.Engine = &gin.Engine{}
)

var (
	Database       *sql.DB = &sql.DB{}
	DatabaseUser   string  = ""
	DatabasePass   string  = ""
	DatabaseHost   string  = ""
	DatabaseName   string  = ""
	DatabasePort   string  = ""
	DatabaseSSL    string  = ""
	DatabaseURL    string  = ""
	DatabaseDriver string  = "postgres"
)

var (
	RabbitURL       string           = ""
	RabbitPort      string           = ""
	RabbitQueueName string           = "queue_name"
	RabbitConn      *amqp.Connection = &amqp.Connection{}
	RabbitChannel   *amqp.Channel    = &amqp.Channel{}
	RabbitQueue     *amqp.Queue      = &amqp.Queue{}
)

var (
	GrpcServer *grpc.Server     = &grpc.Server{}
	GrpcConn   *grpc.ClientConn = &grpc.ClientConn{}
)

var (
	AccountServiceClient  = pb.NewAccountServiceClient(GrpcConn)
	AuthServiceClient     = pb.NewAuthServiceClient(GrpcConn)
	ExerciseServiceClient = pb.NewExerciseServiceClient(GrpcConn)
	ReportServiceClient   = pb.NewReportServiceClient(GrpcConn)
	WorkoutServiceClient  = pb.NewWorkoutServiceClient(GrpcConn)
)
