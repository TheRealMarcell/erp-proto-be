package model

import "fmt"

type InventoryRequest struct {
	Source string							`json:"source"`
	Destination string				`json:"destination"`
	StorageItem
}

func (item InventoryRequest) MoveItem() error{
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
	source_inventory_item.Description = item.Description

	err = source_inventory_item.UpdateItem()
	if err !=nil{
		fmt.Println(err)
		return err
	}

	var destination_inventory_item StorageItem

	destination_inventory_item.Location = item.Destination
	destination_inventory_item.Quantity = item.Quantity
	destination_inventory_item.ItemID = item.ItemID
	destination_inventory_item.Description = item.Description

	err = destination_inventory_item.Save()
	if err !=nil{
		fmt.Println(err)
		return err
	}
	
	return nil
}