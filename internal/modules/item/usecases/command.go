package usecases

import (
	"context"
	"erp-api/internal/modules/item"
	"erp-api/internal/modules/item/models/request"
	"erp-api/internal/pkg/log"
)

type commandUsecase struct {
	itemRepositoryCommand item.PostgresRepositoryCommand
	logger                log.Logger
}

func NewCommandUsecase(
	prq item.PostgresRepositoryCommand,
	log log.Logger) item.UsecaseCommand {
	return commandUsecase{
		itemRepositoryCommand: prq,
		logger:                log,
	}
}

func (c commandUsecase) SaveItem(ctx context.Context, payload request.SubmitItem) error {
	err := c.itemRepositoryCommand.InsertItems(ctx, payload.Items)
	if err != nil {
		return err
	}
	return nil
}
