package main

import (
	"fmt"

	"github.com/username/.../backend/common"
	"github.com/username/.../backend/internal/api/some_service"
	"github.com/username/.../backend/internal/grpc"
	"github.com/username/.../backend/internal/persistence"
	"github.com/username/.../backend/internal/rabbit"
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
	
	some_service.DeclareSomeServiceRoutes(router)

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