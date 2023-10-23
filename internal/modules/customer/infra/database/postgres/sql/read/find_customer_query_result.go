package read

import "time"

type FindCustomerQueryResult struct {
	Document       string    `db:"cpf"`
	Name           string    `db:"name"`
	Phone          string    `db:"phone"`
	Email          string    `db:"email"`
	OptInPromotion bool      `db:"opt_in_promotion"`
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
}
