package request

type itemObject struct {
	ItemID   string `json:"item_id"`
	Quantity int64  `json:"quantity"`
}

type MoveInventory struct {
	Source      string       `json:"source"`
	Destination string       `json:"destination"`
	Items       []itemObject `json:"items"`
}
