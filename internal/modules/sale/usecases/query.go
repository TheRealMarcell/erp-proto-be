package usecases

import (
	"context"
	"erp-api/internal/modules/sale"
	"erp-api/internal/modules/sale/models/entity"
	"erp-api/internal/modules/sale/models/response"
	"erp-api/internal/pkg/errors"
	"erp-api/internal/pkg/log"
)

type queryUsecase struct {
	saleRepositoryQuery sale.PostgresRepositoryQuery
	logger              log.Logger
}

func NewQueryUsecase(
	prq sale.PostgresRepositoryQuery,
	log log.Logger) sale.UsecaseQuery {
	return queryUsecase{
		saleRepositoryQuery: prq,
		logger:              log,
	}
}

func (q queryUsecase) GetSales(ctx context.Context) ([]response.GetSaleResponse, error) {
	respSales := <-q.saleRepositoryQuery.FindAllSales(ctx)
	if respSales.Error != nil {
		msg := "No sales found"
		return nil, errors.NotFound(msg)
	}

	salesData, ok := respSales.Data.([]entity.Sale)

	if !ok {
		return nil, errors.InternalServerError("Cannot parse sales data")
	}

	var resp []response.GetSaleResponse
	for _, sale := range salesData {
		sale := response.GetSaleResponse{
			SaleID:          sale.SaleID,
			ItemID:          sale.ItemID,
			Description:     sale.Description,
			Quantity:        sale.Quantity,
			Price:           sale.Price,
			Total:           sale.Total,
			DiscountPerItem: sale.DiscountPerItem,
			QuantityRetur:   sale.QuantityRetur,
			TransactionID:   sale.TransactionID,
			Location:        sale.Location,
		}
		resp = append(resp, sale)
	}

	return resp, nil
}
