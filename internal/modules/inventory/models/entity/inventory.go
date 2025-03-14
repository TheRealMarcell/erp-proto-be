package entity

type InventoryItem struct {
	ID          int    `json:"id"`
	ItemID      string `json:"item_id"`
	Quantity    int64  `json:"quantity"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
}
