package database

import (
	"context"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
	"hamburgueria/internal/modules/order/domain"
	"hamburgueria/internal/modules/order/domain/valueobject"
	"hamburgueria/internal/modules/order/infra/database/model"
	"hamburgueria/internal/modules/payment/usecase/result"
	"hamburgueria/pkg/sql"
	"sync"
)

type OrderRepository struct {
	readWriteClient *gorm.DB
	readOnlyClient  *gorm.DB
	logger          zerolog.Logger
}

func (c OrderRepository) Create(ctx context.Context, order domain.Order) error {
	err := c.readWriteClient.
		Create(model.FromDomain(order)).Error
	if err != nil {
		c.logger.Error().
			Ctx(ctx).
			Err(err).
			Msg("Failed to insert order")
		return err
	}
	return nil
}

func (c OrderRepository) FindAll(ctx context.Context) ([]domain.Order, error) {
	var orders []model.Order
	tx := c.readOnlyClient.
		Preload("Products").
		Find(&orders)
	if tx.Error != nil {
		c.logger.Error().
			Ctx(ctx).
			Err(tx.Error).
			Msg("Failed to find all orders")
		return nil, tx.Error
	}

	var domainOrders []domain.Order
	for _, order := range orders {
		domainOrders = append(domainOrders, *order.ToDomain())
	}

	return domainOrders, nil
}

func (c OrderRepository) FindByStatus(ctx context.Context, status string) ([]domain.Order, error) {
	var orders []model.Order
	tx := c.readOnlyClient.
		Preload("Products").
		Where("status = ?", status).
		Find(&orders)
	if tx.Error != nil {
		c.logger.Error().
			Ctx(ctx).
			Err(tx.Error).
			Msg("Failed to find orders by status")
		return nil, tx.Error
	}

	var domainOrders []domain.Order
	for _, order := range orders {
		domainOrders = append(domainOrders, *order.ToDomain())
	}

	return domainOrders, nil
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
	readWriteClient *gorm.DB,
	readOnlyClient *gorm.DB,
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
