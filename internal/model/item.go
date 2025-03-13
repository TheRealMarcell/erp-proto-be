package model

import (
	"context"
	db "erp-api/database"
	"fmt"
)

type Item struct {
	ItemID      string `json:"item_id"`
	Price       int64  `json:"price"`
	Description string `json:"description"`
	Quantity    int64  `json:"quantity"`
}

type StorageItem struct {
	Location    string `json:"location"`
	ItemID      string `json:"item_id"`
	Quantity    int64  `json:"quantity"`
	Description string `json:"description"`
}

type ItemList struct {
	Items []Item `json:"items"`
}

type UpdateItemRequest struct {
	Items []UpdateItemObject
}

type UpdateItemObject struct {
	ItemID   string `json:"item_id"`
	Quantity int64  `json:"quantity"`
	SaleID   int64  `json:"sale_id"`
}

type BrokenItem struct {
	ItemID   string `json:"item_id"`
	Quantity int64  `json:"quantity"`
	SaleID   int64  `json:"sale_id"`
}

type BrokenItemRequest struct {
	Items []BrokenItem `json:"items"`
}

type ItemPriceRequest struct {
	ItemID string `json:"item_id"`
	Price  int64  `json:"price"`
}

func GetItemQty(item_id string, location string) (*int64, error) {
	var item_qty int64

	query := fmt.Sprintf(`
	SELECT quantity
	FROM %s
	WHERE item_id = $1`, location)

	err := db.DB.QueryRow(context.Background(), query, item_id).Scan(&item_qty)
	if err != nil {
		return nil, err
	}

	return &item_qty, nil
}

func (item *StorageItem) UpdateItem(operation string) error {
	if operation == "add" {
		qty, err := GetItemQty(item.ItemID, item.Location)

		if err != nil {
			return err
		}

		item.Quantity = item.Quantity + *qty
	}

	query :=
		fmt.Sprintf(`UPDATE %s
	SET quantity = $1
	WHERE item_id = $2`, item.Location)

	_, err := db.DB.Query(context.Background(), query, item.Quantity, item.ItemID)

	if err != nil {
		return err
	}

	return nil
}

func SaveBrokenItem(broken_item BrokenItem) error {
	qty, err := GetItemQty(broken_item.ItemID, "inventory_rusak")
	if err != nil {
		return err
	}

	new_quantity := *qty + broken_item.Quantity

	query := `
	UPDATE inventory_rusak
	SET quantity = $1
	WHERE item_id = $2
	`

	_, err = db.DB.Query(context.Background(), query, new_quantity, broken_item.ItemID)
	if err != nil {
		return err
	}

	query = `
	UPDATE sales
	SET quantity_retur = quantity_retur + $1
	WHERE sale_id = $2`

	_, err = db.DB.Query(context.Background(), query, broken_item.Quantity, broken_item.SaleID)
	if err != nil {
		return err
	}

	return nil

}

func UpdatePrice(item_price ItemPriceRequest) error {
	query := `
	UPDATE items
	SET price = $1
	WHERE item_id = $2
	`

	_, err := db.DB.Query(context.Background(), query, item_price.Price, item_price.ItemID)
	if err != nil {
		return err
	}

	return nil
}
