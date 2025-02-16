package routes

import (
	"encoding/json"
	"erp-api/internal/model"
	"erp-api/util/httpres"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)


func getAllItems(ctx *gin.Context){
	items, err := model.GetItems()
	if err != nil{
		fmt.Println(err)
		httpres.APIResponse(ctx, http.StatusBadRequest, "could not parse request data", nil)
		return
	}
	httpres.APIResponse(ctx, http.StatusOK, "success", items)
}

func insertItem(ctx *gin.Context){
	id := ctx.Param("id")

	var item model.StorageItem

	err := json.NewDecoder(ctx.Request.Body).Decode(&item)
	item.ItemID = id

	if err != nil{
		httpres.APIResponse(ctx, http.StatusBadRequest, "could not parse request data", nil)
		return
	}

	err = item.Save()

	if err != nil{
		fmt.Println(err)
		httpres.APIResponse(ctx, http.StatusInternalServerError, "could not save item", nil)
		return
	}

	httpres.APIResponse(ctx, http.StatusOK, "item inserted!", nil)
}

func createItem(ctx *gin.Context){
	var item model.Item

	err := ctx.ShouldBindJSON(&item)

	if err != nil{
		httpres.APIResponse(ctx, http.StatusBadRequest, "could not parse request data", nil)
		return
	}

	err = item.Create()
	if err != nil{
		httpres.APIResponse(ctx, http.StatusInternalServerError, "could not save item", nil)
		return
	}

	httpres.APIResponse(ctx, http.StatusOK, "success", nil)

}

func updateItem(ctx *gin.Context){
	id := ctx.Param("id")

	var item model.StorageItem

	err := json.NewDecoder(ctx.Request.Body).Decode(&item)
	item.ItemID = id

	if err != nil{
		fmt.Println(err)
		httpres.APIResponse(ctx, http.StatusBadRequest, "could not parse request", nil)
		return
	}


	err = item.UpdateItem("add")
	if err != nil{
		fmt.Println(err)
		httpres.APIResponse(ctx, http.StatusInternalServerError, "could not update item", nil)
		return
	}

	httpres.APIResponse(ctx, http.StatusOK, "success", nil)

}