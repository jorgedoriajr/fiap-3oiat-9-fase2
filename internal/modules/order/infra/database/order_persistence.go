package database

import (
	"context"
	"github.com/rs/zerolog"
	"hamburgueria/internal/modules/order/domain/entity"
	"hamburgueria/internal/modules/order/infra/database/postgres/sql/write"
	"hamburgueria/pkg/querymapper"
	"hamburgueria/pkg/sql"
)

type OrderRepository struct {
	ReadWriteClient sql.Client
	Logger          zerolog.Logger
}

func (c OrderRepository) Create(ctx context.Context, order entity.Order) error {

	mapper := write.EntityToInsertOrderQueryMapper(order)
	args := querymapper.GetArrayOfPropertiesFrom(mapper)

	insertCommand := sql.NewCommand(ctx, c.ReadWriteClient, write.InsertOrderRW, args...)
	err := insertCommand.Exec()

	if err != nil {
		c.Logger.Error().
			Err(err).
			Str("orderId", order.Id.String()).
			Msg("Failed to insert order")
		return err
	}

	return nil
}
