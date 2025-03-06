package model

import (
	"context"
	db "erp-api/database"
	"fmt"
)

type InventoryRequest struct {
	Source string										`json:"source"`
	Destination string							`json:"destination"`
	Items []InventoryMoveItem				`json:"items"`
}

type InventoryMoveItem struct {
	Source string							`json:"source"`
	Destination string				`json:"destination"`
	StorageItem
}

type InventoryItem struct {
	ItemID string 					`json:"item_id"`
	Quantity int64 				 	`json:"quantity"`
	Description string 			`json:"description"`
	Price int64							`json:"price"`
}

func (item InventoryMoveItem) MoveInventory() error{
	qty, err := GetItemQty(item.ItemID, item.Source)

	if err != nil{
		fmt.Println(err)
		return err
	}

	source_qty := *qty - item.Quantity

	var source_inventory_item StorageItem
	source_inventory_item.Location = item.Source
	source_inventory_item.Quantity = source_qty
	source_inventory_item.ItemID = item.ItemID

	err = source_inventory_item.UpdateItem("minus")
	if err !=nil{
		fmt.Println(err)
		return err
	}

	var destination_inventory_item StorageItem

	destination_inventory_item.Location = item.Destination
	destination_inventory_item.Quantity = item.Quantity
	destination_inventory_item.ItemID = item.ItemID

	err = destination_inventory_item.UpdateItem("add")
	if err !=nil{
		fmt.Println(err)
		return err
	}
	
	return nil
}

// GetInventory retrieves inventory from a specified location
func GetInventory(location string) ([]InventoryItem, error) {
	var inventoryTable string

	// Determine the table based on location
	switch location {
	case "toko":
		inventoryTable = "inventory_toko"
	case "tiktok":
		inventoryTable = "inventory_tiktok"
	case "gudang":
		inventoryTable = "inventory_gudang"
	case "rusak":
		inventoryTable = "inventory_rusak"
	default:
		return nil, fmt.Errorf("invalid location: %s", location)
	}

	query := fmt.Sprintf(`
	SELECT i.item_id, quantity, i.description, price 
	FROM %s i
	INNER JOIN items ON i.item_id = items.item_id
	`, inventoryTable)

	rows, err := db.DB.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var inventory []InventoryItem

	for rows.Next() {
		var item InventoryItem
		err = rows.Scan(&item.ItemID, &item.Quantity, &item.Description, &item.Price)
		if err != nil {
			return nil, err
		}
		inventory = append(inventory, item)
	}

	return inventory, nil
}