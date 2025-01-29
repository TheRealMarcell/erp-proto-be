package routes

import (
	"fmt"
	"net/http"

	"erp-api/internal/model"

	"github.com/gin-gonic/gin"
)

// verifyUserByPassword godoc
//
// @Summary 			search user
// @Description 	get user by username, verify if password is correct
// @Tags 					users
// @Accept 				json
// @Produce 			json
// @Param 				json	query	model.UserRequest	true	"Username and Password"
// @Success				200	{array}		model.User
// @Failure 			400 {object}	model.HTTPError
// @Failure 			401 {object}	model.HTTPError
// @Failure 			500 {object}	model.HTTPError
// @Router 				/api/user [get]
func verifyUserByPassword(ctx *gin.Context){
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