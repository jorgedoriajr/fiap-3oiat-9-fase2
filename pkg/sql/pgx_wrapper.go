package sql

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type Rows interface {
	pgx.Rows
}
type Row interface {
	pgx.Row
}
type CommandTag struct {
	ct pgconn.CommandTag
}

func (c CommandTag) String() string {
	return c.ct.String()
}

func (c CommandTag) RowsAffected() int64 {
	return c.ct.RowsAffected()
}

func (c sqlClient) Exec(ctx context.Context, query string, args ...interface{}) (CommandTag, error) {
	ct, err := c.conn.Exec(ctx, query, args...)
	return CommandTag{ct}, err
}

func (c sqlClient) Query(ctx context.Context, query string, args ...interface{}) (Rows, error) {
	rows, err := c.conn.Query(ctx, query, args...)
	return rows, err
}

func (c sqlClient) QueryRow(ctx context.Context, query string, args ...interface{}) Row {
	row := c.conn.QueryRow(ctx, query, args...)
	return row
}

func (c sqlClient) Ping(ctx context.Context) error {
	err := c.conn.Ping(ctx)
	return err
}

func (c sqlClient) SendBatch(ctx context.Context, batch *Batch) (CommandTag, error) {
	br := c.conn.SendBatch(ctx, batch.pgxBatch)
	defer br.Close()
	ct, err := br.Exec()
	return CommandTag{ct}, err
}

func (b *Batch) Queue(query string, args ...any) {
	b.pgxBatch.Queue(query, args...)
}
