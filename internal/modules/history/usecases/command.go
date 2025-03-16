package usecases

import (
	"context"
	"erp-api/internal/modules/history"
	"erp-api/internal/modules/history/models/entity"
	"erp-api/internal/pkg/log"
)

type commandUsecase struct {
	historyRepositoryCommand history.PostgresRepositoryCommand
	logger                   log.Logger
}

func NewCommandUsecase(prc history.PostgresRepositoryCommand, log log.Logger) history.UsecaseCommand {
	return commandUsecase{
		historyRepositoryCommand: prc,
		logger:                   log,
	}
}

func (c commandUsecase) SaveHistory(ctx context.Context, history []entity.History) error {
	return nil
}
