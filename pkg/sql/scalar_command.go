package sql

import (
	"context"

	"github.com/georgysavva/scany/v2/pgxscan"
)

type ScalarCommand[T any] struct {
	command Command
}

func NewScalarCommand[T any](ctx context.Context, client Client, query string, args ...any) *ScalarCommand[T] {
	return &ScalarCommand[T]{command: *NewCommand(ctx, client, query, args...)}
}

func (c ScalarCommand[T]) Exec() (*T, error) {

	var result *T
	err := pgxscan.Get(c.command.ctx, c.command.client.conn, &result, c.command.queryString, c.command.args...)
	if err != nil {
		return nil, err
	}

	return result, nil
}
