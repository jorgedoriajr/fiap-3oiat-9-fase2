package database

import (
	"context"
	"github.com/rs/zerolog"
	"hamburgueria/internal/modules/order/domain"
	"hamburgueria/internal/modules/order/infra/database/postgres/sql/write"
	"hamburgueria/pkg/querymapper"
	"hamburgueria/pkg/sql"
	"sync"
)

type OrderProductRepository struct {
	ReadWriteClient sql.Client
	Logger          zerolog.Logger
}

func (c OrderProductRepository) Create(ctx context.Context, orderProduct domain.OrderProduct) error {

	mapper := write.EntityToInsertOrderProductQueryMapper(orderProduct)
	args := querymapper.GetArrayOfPropertiesFrom(mapper)

	insertCommand := sql.NewCommand(ctx, c.ReadWriteClient, write.InsertOrderProductRW, args...)
	err := insertCommand.Exec()

	if err != nil {
		c.Logger.Error().
			Err(err).
			Str("orderId", orderProduct.OrderId.String()).
			Str("productId", orderProduct.ProductId.String()).
			Msg("Failed to insert product for a order")
		return err
	}

	return nil
}

var (
	orderProductRepositoryInstance OrderProductRepository
	orderProductRepositoryOnce     sync.Once
)

func GetOrderProductPersistence(
	ReadWriteClient sql.Client,
	Logger zerolog.Logger,
) OrderProductRepository {
	orderProductRepositoryOnce.Do(func() {
		orderProductRepositoryInstance = OrderProductRepository{
			ReadWriteClient: ReadWriteClient,
			Logger:          Logger,
		}
	})
	return orderProductRepositoryInstance
}
