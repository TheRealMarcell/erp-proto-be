package response

type Item struct {
	ItemID      string `json:"item_id"`
	Price       int64  `json:"price"`
	Description string `json:"description"`
	Quantity    int64  `json:"quantity"`
}
