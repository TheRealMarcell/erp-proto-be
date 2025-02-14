package model

import (
	"context"
	db "erp-api/database"
	"fmt"
)

type Item struct{
	ItemID string					`json:"item_id"`
	Price int64						`json:"price"`
	Description string		`json:"description"`
}

type StorageItem struct{
	Location string				`json:"location"`
	ItemID string 				`json:"item_id"`
	Quantity int64				`json:"quantity"`
	Description string		`json:"description"`
}

func GetItems() ([]Item, error){
	var items []Item

	query := `
		SELECT * 
		FROM items
	`

	rows, err := db.Conn.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	for rows.Next(){
		var item Item
		err = rows.Scan(&item.Price, &item.Description, &item.ItemID)
		if err != nil{
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}

func GetItemQty(item_id string, location string) (*int64, error){
	var item_qty int64

	query := fmt.Sprintf(`
	SELECT quantity
	FROM %s
	WHERE item_id = $1`, location)

	err := db.Conn.QueryRow(context.Background(), query, item_id).Scan(&item_qty)
	if err != nil{
		return nil, err
	}

	return &item_qty, nil
}

func (item *StorageItem) Save() error{
	query := fmt.Sprintf(
	`INSERT INTO %s (item_id, quantity, description)
	VALUES ($1, $2, $3)`, item.Location)

	fmt.Println(query)

	_, err := db.Conn.Query(context.Background(), query, item.ItemID, item.Quantity, item.Description)

	if err != nil{
		return err
	}

	return nil
}

func (item *StorageItem) UpdateItem() error{
	query := fmt.Sprintf(`UPDATE %s
	SET quantity = $1
	WHERE item_id = $2`, item.Location)

	_, err := db.Conn.Query(context.Background(), query, item.Quantity, item.ItemID)

	if err != nil{
		return err
	}

	return nil
}