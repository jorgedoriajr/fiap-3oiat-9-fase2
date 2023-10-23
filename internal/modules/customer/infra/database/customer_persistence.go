package database

import (
	"context"
	"github.com/rs/zerolog"
	"hamburgueria/internal/modules/customer/domain/entity"
	"hamburgueria/internal/modules/customer/infra/database/postgres/sql/read"
	"hamburgueria/internal/modules/customer/infra/database/postgres/sql/write"
	queryhelper "hamburgueria/pkg/querymapper"
	"hamburgueria/pkg/sql"
)

type CustomerRepository struct {
	ReadWriteClient sql.Client
	ReadOnlyClient  sql.Client
	Logger          zerolog.Logger
}

func (c CustomerRepository) Create(ctx context.Context, customer entity.Customer) error {

	mapper := write.InsertCustomerRWQueryMapper{
		Document:       customer.Document,
		Name:           customer.Name,
		Phone:          customer.Phone,
		Email:          customer.Email,
		OptInPromotion: customer.OptInPromotion,
		CreatedAt:      customer.CreatedAt,
		UpdatedAt:      customer.UpdatedAt,
	}
	args := queryhelper.GetArrayOfPropertiesFrom(mapper)

	insertCommand := sql.NewCommand(ctx, c.ReadWriteClient, write.InsertCustomerRW, args...)

	err := insertCommand.Exec()

	if err != nil {
		c.Logger.Error().
			Err(err).
			Str("document", customer.Document).
			Msg("Failed to insert customer")
		return err
	}

	return nil
}

func (c CustomerRepository) Get(ctx context.Context, document string) (customerResult *read.FindCustomerQueryResult, err error) {

	row, err := sql.NewQuery[*read.FindCustomerQueryResult](ctx, c.ReadOnlyClient, read.FindCustomerByCpf, document).One()

	if err != nil {
		c.Logger.Error().
			Err(err).
			Str("document", document).
			Msg("Failed to get customer")
		return nil, err
	}

	return row, nil
}
