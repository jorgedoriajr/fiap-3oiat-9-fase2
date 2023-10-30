package sql

import (
	"context"
	"github.com/georgysavva/scany/v2/pgxscan"
)

type Query[T any] struct {
	command Command
}

func NewQuery[T any](ctx context.Context, client Client, query string, args ...any) *Query[T] {
	return &Query[T]{command: *NewCommand(ctx, client, query, args...)}
}

func (q Query[T]) One() (T, error) {
	//TODO: metric

	var zeroValue T

	result, err := q.Many()
	if err != nil {
		return zeroValue, err
	}

	if len(result) > 0 {
		return result[0], nil
	}

	return zeroValue, nil
}

func (q Query[T]) Many() ([]T, error) {
	var result []T

	err := pgxscan.Select(q.command.ctx, q.command.client.conn, &result, q.command.queryString, q.command.args...)
	return result, err
}

func (q Query[T]) Insert() Row {
	return q.command.client.QueryRow(q.command.ctx, q.command.queryString, q.command.args...)
}
