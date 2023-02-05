package main

import (
	"fmt"

	"github.com/carlosabdoamaral/behealthy_api/common"
	auth_api "github.com/carlosabdoamaral/behealthy_api/internal/api/auth"
	"github.com/carlosabdoamaral/behealthy_api/internal/grpc"
	"github.com/carlosabdoamaral/behealthy_api/internal/persistence"
	"github.com/carlosabdoamaral/behealthy_api/internal/rabbit"
	"github.com/gin-gonic/gin"
)

func main() {
	common.GetEnvVariables()

	_, err := persistence.Connect()
	if err != nil {
		fmt.Printf("[!] %s", err.Error())
		return
	}

	rabbit.Connect()
	grpc.ConnectToGRPCServer()

	router := gin.Default()
	router.Use(CORS())

	routesGroup := router.Group("/api")
	auth_api.DeclareAuthRoutes(routesGroup)

	router.Run()
}

// https://github.com/gin-contrib/cors/issues/29
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
