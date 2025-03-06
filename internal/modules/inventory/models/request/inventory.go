package request

type InventoryMoveItemRequest struct {
	Source string							`json:"source"`
	Destination string				`json:"destination"`
	Location string				`json:"location"`
	ItemID string 				`json:"item_id"`
	Quantity int64				`json:"quantity"`
	Description string		`json:"description"`
}

type GetInventoryRequest struct {
	Source string														`json:"source"`
	Destination string											`json:"destination"`
	Items []InventoryMoveItemRequest				`json:"items"`
}