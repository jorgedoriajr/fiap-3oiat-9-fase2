package database

import (
	"context"
	"github.com/rs/zerolog"
	"hamburgueria/internal/modules/order/domain/entity"
	"hamburgueria/internal/modules/order/infra/database/postgres/sql/read"
	"hamburgueria/internal/modules/order/infra/database/postgres/sql/write"
	"hamburgueria/pkg/querymapper"
	"hamburgueria/pkg/sql"
	"sync"
)

type OrderRepository struct {
	readWriteClient sql.Client
	readOnlyClient  sql.Client
	logger          zerolog.Logger
}

func (c OrderRepository) Create(ctx context.Context, order entity.Order) error {

	mapper := write.EntityToInsertOrderQueryMapper(order)
	args := querymapper.GetArrayOfPropertiesFrom(mapper)

	insertCommand := sql.NewCommand(ctx, c.readWriteClient, write.InsertOrderRW, args...)
	err := insertCommand.Exec()

	if err != nil {
		c.logger.Error().
			Err(err).
			Str("orderId", order.Id.String()).
			Msg("Failed to insert order")
		return err
	}

	return nil
}

func (c OrderRepository) FindAll(ctx context.Context) ([]entity.Order, error) {
	allOrders, err := sql.NewQuery[read.FindOrderQueryResult](
		ctx,
		c.readOnlyClient,
		read.FindAllOrders,
	).Many()

	if err != nil {
		c.logger.Error().
			Err(err).
			Msg("Failed to get orders")
		return nil, err
	}

	var orders []entity.Order
	for _, order := range allOrders {
		orders = append(orders, order.ToEntity())
	}
	return orders, nil
}

func (c OrderRepository) FindByStatus(ctx context.Context, status string) ([]entity.Order, error) {
	ordersByStatus, err := sql.NewQuery[read.FindOrderQueryResult](
		ctx,
		c.readOnlyClient,
		read.FindOrderByStatus,
		status,
	).Many()

	if err != nil {
		c.logger.Error().
			Err(err).
			Msg("Failed to get orders by status")
		return nil, err
	}

	var orders []entity.Order
	for _, order := range ordersByStatus {
		orders = append(orders, order.ToEntity())
	}
	return orders, nil
}

var (
	orderRepositoryInstance OrderRepository
	orderRepositoryOnce     sync.Once
)

func GetOrderPersistence(
	readWriteClient sql.Client,
	readOnlyClient sql.Client,
	logger zerolog.Logger,
) OrderRepository {
	orderRepositoryOnce.Do(func() {
		orderRepositoryInstance = OrderRepository{
			readWriteClient: readWriteClient,
			readOnlyClient:  readOnlyClient,
			logger:          logger,
		}
	})
	return orderRepositoryInstance
}
