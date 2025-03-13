package request

import "erp-api/internal/modules/item/models/entity"

type SubmitItem struct {
	Items []entity.Item `json:"items"`
}
