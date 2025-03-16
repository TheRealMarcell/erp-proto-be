package response

import "time"

type History struct {
	PindahanID  int64     `json:"pindahan_id"`
	ItemID      string    `json:"item_id"`
	Quantity    int64     `json:"quantity"`
	Timestamp   time.Time `json:"timestamp"`
	Source      string    `json:"source"`
	Destination string    `json:"destination"`
}
