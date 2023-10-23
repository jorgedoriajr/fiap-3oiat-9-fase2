package sql

import (
	"context"
	"fmt"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/ext"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"

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
	span, ctx := tracer.StartSpanFromContext(
		ctx,
		"postgres.operations",
		tracer.ServiceName(c.applicationName),
		tracer.ResourceName(fmt.Sprintf("Exec %s", query)),
		tracer.SpanType(ext.SpanTypeSQL),
		tracer.Tag(ext.Component, "ifood/go-packages"),
		tracer.Tag(ext.SpanKind, ext.SpanKindClient),
		tracer.Tag(ext.DBType, ext.DBSystemPostgreSQL),
		tracer.Measured(),
	)
	ct, err := c.conn.Exec(ctx, query, args...)
	span.Finish(tracer.WithError(err))
	return CommandTag{ct}, err
}

func (c sqlClient) Query(ctx context.Context, query string, args ...interface{}) (Rows, error) {
	span, ctx := tracer.StartSpanFromContext(
		ctx,
		"postgres.operations",
		tracer.ServiceName(c.applicationName),
		tracer.ResourceName(fmt.Sprintf("Query %s", query)),
		tracer.SpanType(ext.SpanTypeSQL),
		tracer.Tag(ext.Component, "ifood/go-packages"),
		tracer.Tag(ext.SpanKind, ext.SpanKindClient),
		tracer.Tag(ext.DBType, ext.DBSystemPostgreSQL),
		tracer.Measured(),
	)
	rows, err := c.conn.Query(ctx, query, args...)
	span.Finish(tracer.WithError(err))
	return rows, err
}

func (c sqlClient) QueryRow(ctx context.Context, query string, args ...interface{}) Row {
	span, ctx := tracer.StartSpanFromContext(
		ctx,
		"postgres.operations",
		tracer.ServiceName(c.applicationName),
		tracer.ResourceName(fmt.Sprintf("QueryRow %s", query)),
		tracer.SpanType(ext.SpanTypeSQL),
		tracer.Tag(ext.Component, "ifood/go-packages"),
		tracer.Tag(ext.SpanKind, ext.SpanKindClient),
		tracer.Tag(ext.DBType, ext.DBSystemPostgreSQL),
		tracer.Measured(),
	)
	row := c.conn.QueryRow(ctx, query, args...)
	span.Finish()
	return row
}

func (c sqlClient) Ping(ctx context.Context) error {
	span, ctx := tracer.StartSpanFromContext(
		ctx,
		"postgres.operations",
		tracer.ServiceName(c.applicationName),
		tracer.ResourceName("Ping"),
		tracer.SpanType(ext.SpanTypeSQL),
		tracer.Tag(ext.Component, "ifood/go-packages"),
		tracer.Tag(ext.SpanKind, ext.SpanKindClient),
		tracer.Tag(ext.DBType, ext.DBSystemPostgreSQL),
		tracer.Measured(),
	)
	err := c.conn.Ping(ctx)
	span.Finish(tracer.WithError(err))
	return err
}

func (c sqlClient) SendBatch(ctx context.Context, batch *Batch) (CommandTag, error) {
	span, ctx := tracer.StartSpanFromContext(
		ctx,
		"postgres.operations",
		tracer.ServiceName(c.applicationName),
		tracer.ResourceName("SendBatch"),
		tracer.SpanType(ext.SpanTypeSQL),
		tracer.Tag(ext.Component, "ifood/go-packages"),
		tracer.Tag(ext.SpanKind, ext.SpanKindClient),
		tracer.Tag(ext.DBType, ext.DBSystemPostgreSQL),
		tracer.Measured(),
	)
	br := c.conn.SendBatch(ctx, batch.pgxBatch)
	defer br.Close()
	ct, err := br.Exec()
	span.Finish(tracer.WithError(err))
	return CommandTag{ct}, err
}

func (b *Batch) Queue(query string, args ...any) {
	b.pgxBatch.Queue(query, args...)
}
