package routes

import (
	"fmt"
	"net/http"

	"erp-api/internal/model"
	"erp-api/internal/pkg/util/httpres"

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
func verifyUserByPassword(ctx *gin.Context) {
	var userReq model.UserRequest

	err := ctx.ShouldBindJSON(&userReq)
	if err != nil {
		fmt.Println(err)
		httpres.APIResponse(ctx, http.StatusBadRequest, "failed to parse request data", nil)
		return
	}

	// get user from database, pass the req body
	user, err := model.GetUser(userReq)
	if err != nil {
		httpres.APIResponse(ctx, http.StatusInternalServerError, "could not fetch user", nil)
		return
	}

	err = model.VerifyUser(userReq, user)
	if err != nil {
		httpres.APIResponse(ctx, http.StatusUnauthorized, "unauthorised login, wrong password", nil)
		return
	}
	httpres.APIResponse(ctx, http.StatusOK, "successfully verified user", user)
}
