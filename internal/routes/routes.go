package routes

import (
	"github.com/gin-gonic/gin"
)


func RegisterRoutes(server *gin.RouterGroup){
	server.GET("/user", getUserByUsername)

}