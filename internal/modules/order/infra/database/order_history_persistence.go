package database

import (
	"context"
	"github.com/rs/zerolog"
	"hamburgueria/internal/modules/order/domain/entity"
	"hamburgueria/internal/modules/order/infra/database/postgres/sql/write"
	"hamburgueria/pkg/querymapper"
	"hamburgueria/pkg/sql"
	"sync"
)

type OrderHistoryRepository struct {
	ReadWriteClient sql.Client
	Logger          zerolog.Logger
}

func (c OrderHistoryRepository) Create(ctx context.Context, orderHistory entity.OrderHistory) error {

	mapper := write.EntityToInsertOrderHistoryQueryMapper(orderHistory)
	args := querymapper.GetArrayOfPropertiesFrom(mapper)

	insertCommand := sql.NewCommand(ctx, c.ReadWriteClient, write.InsertOrderHistoryRW, args...)
	err := insertCommand.Exec()

	if err != nil {
		c.Logger.Error().
			Err(err).
			Str("orderId", orderHistory.OrderId.String()).
			Msg("Failed to insert order history event")
		return err
	}

	return nil
}

var (
	orderHistoryRepositoryInstance OrderHistoryRepository
	orderHistoryRepositoryOnce     sync.Once
)

func GetOrderHistoryPersistence(
	ReadWriteClient sql.Client,
	Logger zerolog.Logger,
) OrderHistoryRepository {
	orderHistoryRepositoryOnce.Do(func() {
		orderHistoryRepositoryInstance = OrderHistoryRepository{
			ReadWriteClient: ReadWriteClient,
			Logger:          Logger,
		}
	})
	return orderHistoryRepositoryInstance
}
