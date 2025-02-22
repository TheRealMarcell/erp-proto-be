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
		fmt.Println(err)
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
	var itemreq model.ItemList

	err := ctx.ShouldBindJSON(&itemreq)

	if err != nil{
		fmt.Println(err)
		httpres.APIResponse(ctx, http.StatusBadRequest, "could not parse request", nil)
		return
	}

	for _, item := range itemreq.Items{
		err = item.Create()

		if err != nil{
			fmt.Println(err)
			httpres.APIResponse(ctx, http.StatusBadRequest, "could not parse save item", nil)
			return
		}
	
		var storageitem model.StorageItem
		storageitem.ItemID = item.ItemID
		storageitem.Description = item.Description
		storageitem.Location = "inventory_gudang"
		storageitem.Quantity = item.Quantity
	
		err = storageitem.Save()
	
		if err != nil{
			fmt.Println(err)
			httpres.APIResponse(ctx, http.StatusInternalServerError, "could not save item", nil)
			return
		}

	}

	httpres.APIResponse(ctx, http.StatusOK, "success", nil)

}

func updateItem(ctx *gin.Context){
	var update_request model.UpdateItemRequest

	err := ctx.ShouldBindJSON(&update_request)
	if err != nil{
		fmt.Println(err)
		httpres.APIResponse(ctx, http.StatusBadRequest, "could not parse request", nil)
		return
	}

	for _, item := range update_request.Items{
		var storage_item model.StorageItem
		storage_item.ItemID = item.ItemID
		storage_item.Quantity = item.Quantity
		storage_item.Location = "inventory_gudang"

		err = storage_item.UpdateItem("add")
		if err != nil{
			fmt.Println(err)
			httpres.APIResponse(ctx, http.StatusInternalServerError, "could not update item", nil)
			return
		}	

		err = model.UpdateSaleReturQty(item.Quantity, item.SaleID)
		if err != nil{
			fmt.Println(err)
			httpres.APIResponse(ctx, http.StatusInternalServerError, "could not update sale retur qty", nil)
			return
		}	
	}

	httpres.APIResponse(ctx, http.StatusOK, "success", nil)
}


func correctItem(ctx *gin.Context){
	id := ctx.Param("id")

	var item model.StorageItem

	err := json.NewDecoder(ctx.Request.Body).Decode(&item)
	item.ItemID = id

	if err != nil{
		fmt.Println(err)
		httpres.APIResponse(ctx, http.StatusBadRequest, "could not parse request", nil)
		return
	}

	err = item.UpdateItem("set")
	if err != nil{
		fmt.Println(err)
		httpres.APIResponse(ctx, http.StatusInternalServerError, "could not update item", nil)
		return
	}

	httpres.APIResponse(ctx, http.StatusOK, "success", nil)
}

func brokenItem(ctx *gin.Context){
	// update total on sale id, insert to inventory_rusak

	var broken_item model.BrokenItemRequest

	err := ctx.ShouldBindJSON(&broken_item)
	if err != nil{
		httpres.APIResponse(ctx, http.StatusBadRequest, "could not parse request", nil)
		return
	}

	for _, broken_item := range broken_item.Items{
		err = model.SaveBrokenItem(broken_item)
		if err != nil{
			httpres.APIResponse(ctx, http.StatusBadRequest, "could not update database", nil)
			return
		}
	}	

	httpres.APIResponse(ctx, http.StatusOK, "successfully updated broken item", nil)
}

func updateItemPrice(ctx *gin.Context){
	var item_request model.ItemPriceRequest

	err := ctx.ShouldBindJSON(&item_request)
	if err != nil{
		httpres.APIResponse(ctx, http.StatusBadRequest, "could not parse request", nil)
		return
	}

	err = model.UpdatePrice(item_request)
	if err != nil{
		httpres.APIResponse(ctx, http.StatusInternalServerError, "could not update price", nil)
		return
	}

	httpres.APIResponse(ctx, http.StatusOK, "successfully updated price", nil)

}