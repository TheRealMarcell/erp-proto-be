package model

type InventoryRequest struct {
	Source string							`json:"source"`
	Destination string				`json:"destination"`
	StorageItem
}