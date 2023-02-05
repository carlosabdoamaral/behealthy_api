package auth

import (
	"github.com/gin-gonic/gin"
)

func DeclareAuthRoutes(router *gin.RouterGroup) {
	g := router.Group("/auth")
	v1 := g.Group("/v1")
	v1.POST("/sign-in", HandleSignIn)
	v1.POST("/sign-up", HandleSignUp)
	v1.PUT("/two-fa/refresh", HandleRefreshTwoFactorCode)
}
