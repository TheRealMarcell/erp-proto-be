package routes

import (
	"fmt"
	"net/http"

	"erp-api/internal/model"

	"github.com/gin-gonic/gin"
)

func getUserByUsername(ctx *gin.Context){
	var userReq model.UserRequest

	err := ctx.ShouldBindJSON(&userReq)
	if err != nil{
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "failed to parse request"})
		return
	}

	// get user from database, pass the req body
	user, err := model.GetUser(userReq)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch user"})
		return
	}

	err = model.VerifyUser(userReq, user)
	if err != nil{
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorised login, wrong password"})
		return
	}

	ctx.JSON(http.StatusOK, user)
}