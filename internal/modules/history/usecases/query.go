package usecases

import (
	"context"
	"erp-api/internal/modules/history"
	"erp-api/internal/modules/history/models/entity"
	"erp-api/internal/modules/history/models/response"
	"erp-api/internal/pkg/errors"
	"erp-api/internal/pkg/log"
)

type queryUsecase struct {
	historyRepositoryQuery history.PostgresRepositoryQuery
	logger                 log.Logger
}

func NewQueryUsecase(prq history.PostgresRepositoryQuery, log log.Logger) history.UsecaseQuery {
	return queryUsecase{
		historyRepositoryQuery: prq,
		logger:                 log,
	}
}

func (q queryUsecase) GetHistory(ctx context.Context) ([]response.History, error) {
	respHistory := <-q.historyRepositoryQuery.GetListHistory(ctx)
	if respHistory.Error != nil {
		return nil, errors.NotFound("could not get history")
	}

	historyData, ok := respHistory.Data.([]entity.History)

	if !ok {
		return nil, errors.InternalServerError("cannot parse history")
	}

	var resp []response.History

	for _, hi := range historyData {
		history := response.History{
			PindahanID:  hi.PindahanID,
			ItemID:      hi.ItemID,
			Quantity:    hi.Quantity,
			Timestamp:   hi.Timestamp,
			Source:      hi.Source,
			Destination: hi.Destination,
			GroupID:     hi.GroupID,
		}
		resp = append(resp, history)
	}
	return resp, nil
}
