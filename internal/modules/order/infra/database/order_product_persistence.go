package database

import (
	"context"
	"github.com/rs/zerolog"
	"hamburgueria/internal/modules/order/domain/entity"
	"hamburgueria/internal/modules/order/infra/database/postgres/sql/write"
	"hamburgueria/pkg/querymapper"
	"hamburgueria/pkg/sql"
)

type OrderProductRepository struct {
	ReadWriteClient sql.Client
	Logger          zerolog.Logger
}

func (c OrderProductRepository) Create(ctx context.Context, orderProduct entity.OrderProduct) error {

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
