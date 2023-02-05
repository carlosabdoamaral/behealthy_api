package account

import "github.com/gin-gonic/gin"

func DeclareAccountRoutes(router *gin.RouterGroup) {
	g := router.Group("/account")
	g.PUT("/update-password", HandleUpdatePassword)
	g.PUT("/soft-delete", HandleSoftDelete)
	g.PUT("/restore", HandleRestore)
	g.PUT("/block", HandleBlock)
	g.PUT("/unblock", HandleUnblock)
}
