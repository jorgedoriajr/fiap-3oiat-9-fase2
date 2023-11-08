package sql

import (
	"context"
)

type Command struct {
	queryString string
	args        []any
	ctx         context.Context
	client      *sqlClient
}

func NewCommand(ctx context.Context, client Client, query string, args ...any) *Command {
	return &Command{
		queryString: query,
		args:        args,
		ctx:         ctx,
		client:      client.(*sqlClient),
	}
}

func (c Command) Exec() error {

	_, err := c.client.Exec(c.ctx, c.queryString, c.args...)
	if err != nil {
		return err
	}

	return nil
}
