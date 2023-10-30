package database

import (
	"context"
	"github.com/rs/zerolog"
	"hamburgueria/internal/modules/customer/domain/entity"
	"hamburgueria/internal/modules/customer/infra/database/postgres/sql/read"
	"hamburgueria/internal/modules/customer/infra/database/postgres/sql/write"
	"hamburgueria/pkg/querymapper"
	"hamburgueria/pkg/sql"
	"sync"
)

type CustomerRepository struct {
	readWriteClient sql.Client
	readOnlyClient  sql.Client
	logger          zerolog.Logger
}

func (c CustomerRepository) Create(ctx context.Context, customer entity.Customer) error {

	mapper := write.EntityToInsertCustomerQueryMapper(customer)
	args := querymapper.GetArrayOfPropertiesFrom(mapper)

	insertCommand := sql.NewCommand(ctx, c.readWriteClient, write.InsertCustomerRW, args...)
	err := insertCommand.Exec()

	if err != nil {
		c.logger.Error().
			Err(err).
			Str("document", customer.Document).
			Msg("Failed to insert customer")
		return err
	}

	return nil
}

func (c CustomerRepository) Get(ctx context.Context, document string) (customerResult *entity.Customer, err error) {

	result, err := sql.NewQuery[*read.FindCustomerQueryResult](ctx, c.readOnlyClient, read.FindCustomerByCpf, document).One()

	if err != nil {
		c.logger.Error().
			Err(err).
			Str("document", document).
			Msg("Failed to get customer")
		return nil, err
	}

	return result.ToEntity(), nil
}

var (
	customerRepositoryInstance CustomerRepository
	customerRepositoryOnce     sync.Once
)

func GetCustomerPersistence(
	ReadWriteClient sql.Client,
	ReadOnlyClient sql.Client,
	Logger zerolog.Logger,
) CustomerRepository {
	customerRepositoryOnce.Do(func() {
		customerRepositoryInstance = CustomerRepository{
			readWriteClient: ReadWriteClient,
			readOnlyClient:  ReadOnlyClient,
			logger:          Logger,
		}
	})
	return customerRepositoryInstance
}
