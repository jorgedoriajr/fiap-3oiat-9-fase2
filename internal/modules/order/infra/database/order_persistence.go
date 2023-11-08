package database

import (
	"context"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"hamburgueria/internal/modules/order/domain"
	"hamburgueria/internal/modules/order/domain/valueobject"
	"hamburgueria/internal/modules/order/infra/database/postgres/sql/read"
	"hamburgueria/internal/modules/order/infra/database/postgres/sql/write"
	"hamburgueria/internal/modules/payment/usecase/result"
	"hamburgueria/pkg/querymapper"
	"hamburgueria/pkg/sql"
	"sync"
)

type OrderRepository struct {
	readWriteClient sql.Client
	readOnlyClient  sql.Client
	logger          zerolog.Logger
}

func (c OrderRepository) Create(ctx context.Context, order domain.Order) error {

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

func (c OrderRepository) FindAll(ctx context.Context) ([]domain.Order, error) {
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

	var orders []domain.Order
	for _, order := range allOrders {
		orders = append(orders, order.ToEntity())
	}
	return orders, nil
}

func (c OrderRepository) FindByStatus(ctx context.Context, status string) ([]domain.Order, error) {
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

	var orders []domain.Order
	for _, order := range ordersByStatus {
		orders = append(orders, order.ToEntity())
	}
	return orders, nil
}

func (c OrderRepository) SavePaymentReference(ctx context.Context, payment result.PaymentProcessed) error {
	err := sql.NewCommand(
		ctx,
		c.readWriteClient,
		write.UpdateOrderPayment,
		payment.PaymentId,
		string(valueobject.PaymentCreated),
		payment.OrderReference,
	).Exec()
	if err != nil {
		return err
	}

	return nil
}

func (c OrderRepository) FindById(ctx context.Context, orderId uuid.UUID) (*domain.Order, error) {
	order, err := sql.NewQuery[*read.FindOrderQueryResult](
		ctx,
		c.readOnlyClient,
		read.FindOrderById,
		orderId.String(),
	).One()

	if err != nil {
		c.logger.Error().
			Err(err).
			Msg("Failed to get orders by status")
		return nil, err
	}

	if order == nil {
		return nil, nil
	}
	orderEntity := order.ToEntity()
	return &orderEntity, nil
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
