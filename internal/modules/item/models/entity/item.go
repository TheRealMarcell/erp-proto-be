package entity

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
