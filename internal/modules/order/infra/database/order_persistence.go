package database

import (
	"context"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"hamburgueria/internal/modules/order/domain"
	"hamburgueria/internal/modules/order/infra/database/model"
	"sync"
)

type OrderRepository struct {
	readWriteClient *gorm.DB
	readOnlyClient  *gorm.DB
	logger          zerolog.Logger
}

func (c OrderRepository) Create(ctx context.Context, order domain.Order) error {
	orderModel := model.FromDomain(order)
	orderModel.History = []model.OrderHistory{{
		ID:        uuid.New(),
		OrderId:   orderModel.ID,
		Status:    orderModel.Status,
		ChangeBy:  "SYSTEM",
		CreatedAt: orderModel.CreatedAt,
	}}
	err := c.readWriteClient.
		Create(orderModel).Error
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
	err := c.readOnlyClient.
		Preload(clause.Associations).
		Find(&orders).Error
	if err != nil {
		c.logger.Error().
			Ctx(ctx).
			Err(err).
			Msg("Failed to find all orders")
		return nil, err
	}

	var domainOrders []domain.Order
	for _, order := range orders {
		domainOrders = append(domainOrders, *order.ToDomain())
	}

	return domainOrders, nil
}

func (c OrderRepository) FindByStatus(ctx context.Context, status string) ([]domain.Order, error) {
	var orders []model.Order
	err := c.readOnlyClient.
		Preload(clause.Associations).
		Where("status = ?", status).
		Find(&orders).Error
	if err != nil {
		c.logger.Error().
			Ctx(ctx).
			Err(err).
			Str("status", status).
			Msg("Failed to find orders by status")
		return nil, err
	}

	var domainOrders []domain.Order
	for _, order := range orders {
		domainOrders = append(domainOrders, *order.ToDomain())
	}

	return domainOrders, nil
}

func (c OrderRepository) Update(ctx context.Context, order domain.Order) error {
	orderModel := model.FromDomain(order)
	err := c.readWriteClient.
		Session(&gorm.Session{FullSaveAssociations: true}).
		Save(&orderModel).
		Error
	if err != nil {
		c.logger.Error().
			Ctx(ctx).
			Err(err).
			Str("orderId", order.Id.String()).
			Msg("Failed to update order")
		return err
	}
	return nil
}

func (c OrderRepository) FindById(ctx context.Context, orderId uuid.UUID) (*domain.Order, error) {
	var order model.Order
	tx := c.readOnlyClient.
		Preload("Products").
		Preload("History").
		Find(&order, orderId)
	if tx.Error != nil {
		c.logger.Error().
			Ctx(ctx).
			Err(tx.Error).
			Str("orderId", orderId.String()).
			Msg("Failed to find orders by ID")
		return nil, tx.Error
	}

	return order.ToDomain(), nil
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
