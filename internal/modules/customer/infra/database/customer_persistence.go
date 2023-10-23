package database

import (
	"context"
	"fmt"
	"github.com/rs/zerolog"
	"hamburgueria/internal/modules/customer/domain/entity"
	"hamburgueria/pkg/sql"
)

type CustomerRepository struct {
	ReadWriteClient sql.Client
	ReadOnlyClient  sql.Client
	Logger          zerolog.Logger
}

func (c CustomerRepository) Create(ctx context.Context, customer entity.Customer) error {
	insertStatement := fmt.Sprintf(`
		INSERT INTO customer (
		  cpf, 
		  phone,            
		  name,             
		  email,            
		  opt_in_promotion,       
		  created_at,
		  updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)`,
	)

	insertCommand := sql.NewCommand(
		ctx,
		c.ReadWriteClient,
		insertStatement,
		customer.Document,
		customer.Phone,
		customer.Name,
		customer.Email,
		customer.OptInPromotion,
		customer.CreatedAt,
		customer.UpdatedAt,
	)

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

func (c CustomerRepository) Get(ctx context.Context, document string) (customerResult *entity.Customer, err error) {
	query := fmt.Sprintf(`
		SELECT
			cpf,              
			phone,            
			name,             
			email,            
			opt_in_promotion,       
			created_at,
			updated_at
		FROM customer
		WHERE cpf = $1
		LIMIT 1`,
	)

	row, err := sql.NewQuery[*entity.Customer](ctx, c.ReadOnlyClient, query, document).One()

	if err != nil {
		c.Logger.Error().
			Err(err).
			Str("document", document).
			Msg("Failed to get customer")
		return nil, err
	}

	return row, nil
}
