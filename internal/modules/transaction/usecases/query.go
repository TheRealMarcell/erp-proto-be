package usecases

import (
	"context"
	"erp-api/internal/modules/transaction"
	"erp-api/internal/modules/transaction/models/entity"
	"erp-api/internal/modules/transaction/models/response"
	"erp-api/internal/pkg/errors"
	"erp-api/internal/pkg/log"
)

type queryUsecase struct {
	transactionRepositoryQuery transaction.PostgresRepositoryQuery
	logger                     log.Logger
}

func NewQueryUsecase(prq transaction.PostgresRepositoryQuery, logger log.Logger) transaction.UsecaseQuery {
	return queryUsecase{
		transactionRepositoryQuery: prq,
		logger:                     logger,
	}
}

func (q queryUsecase) GetTransactions(ctx context.Context) ([]response.Transaction, error) {
	respTransactions := <-q.transactionRepositoryQuery.FindAllTransactions(ctx)
	if respTransactions.Error != nil {
		return nil, errors.NotFound("could not fetch transactions")
	}

	transactionData, ok := respTransactions.Data.([]entity.Transaction)

	if !ok {
		return nil, errors.InternalServerError("cannot parse transaction data")
	}

	var resp []response.Transaction

	for _, t := range transactionData {
		transaction := response.Transaction{
			TransactionID:   t.TransactionID,
			DiscountType:    t.DiscountType,
			DiscountPercent: t.DiscountPercent,
			TotalDiscount:   t.TotalDiscount,
			TotalPrice:      t.TotalPrice,
			PaymentID:       t.PaymentID,
			CustomerName:    t.CustomerName,
			Timestamp:       t.Timestamp,
			Location:        t.Location,
			PaymentStatus:   t.PaymentStatus,
			DownPayment:     t.DownPayment,
		}
		resp = append(resp, transaction)
	}
	return resp, nil
}

func (q queryUsecase) GetDiscountPercentages(ctx context.Context) ([]response.GetDiscount, error) {
	respTransactionDiscounts := <-q.transactionRepositoryQuery.FindDiscount(ctx)
	if respTransactionDiscounts.Error != nil {
		return nil, errors.NotFound("could not fetch transaction discounts")
	}

	discountData, ok := respTransactionDiscounts.Data.([]entity.TransactionDiscount)
	if !ok {
		return nil, errors.InternalServerError("cannot parse transaction discounts data")
	}

	var resp []response.GetDiscount

	for _, d := range discountData {
		discount := response.GetDiscount{
			TransactionID:   d.TransactionID,
			DiscountPercent: d.DiscountPercent,
		}

		resp = append(resp, discount)
	}

	return resp, nil
}
