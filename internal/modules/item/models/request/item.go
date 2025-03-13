package request

import "erp-api/internal/modules/item/models/entity"

type UpdateItemObject struct {
	ItemID   string `json:"item_id"`
	Quantity int64  `json:"quantity"`
	SaleID   int64  `json:"sale_id"`
}

type SubmitItem struct {
	Items []entity.Item `json:"items"`
}

type UpdateItem struct {
	Items []UpdateItemObject `json:"items"`
}
