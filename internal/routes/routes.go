package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.RouterGroup) {
	server.POST("/verify-user", verifyUserByPassword)
}
