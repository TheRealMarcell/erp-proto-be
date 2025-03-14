package request

import "erp-api/internal/modules/item/models/entity"

type updateItemObject struct {
	ItemID   string `json:"item_id"`
	Quantity int64  `json:"quantity"`
	SaleID   int64  `json:"sale_id"`
}

type SubmitItem struct {
	Items []entity.Item `json:"items"`
}

type UpdateItem struct {
	Items []updateItemObject `json:"items"`
}

type BrokenItem struct {
	Items []updateItemObject `json:"items"`
}

type CorrectItem struct {
	Location string `json:"location"`
	Quantity int64  `json:"quantity"`
}

type ItemPrice struct {
	ItemID string `json:"item_id"`
	Price  int64  `json:"price"`
}
