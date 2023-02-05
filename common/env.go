package common

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
	pb "github.com/username/.../backend/protodefs/gen/proto"
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
	GrpcServer              *grpc.Server     = &grpc.Server{}
	GrpcConn                *grpc.ClientConn = &grpc.ClientConn{}
	// AccountServiceClient                     = pb.NewAccountServiceClient(GrpcConn) // Just an example
)
