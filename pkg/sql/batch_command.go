package sql

import "context"

type BatchCommand struct {
	Command
	batch *Batch
}

func NewBatchCommand(ctx context.Context, client Client, query string) *BatchCommand {
	return &BatchCommand{
		Command: *NewCommand(ctx, client, query),
		batch:   NewBatch(),
	}
}

func (c BatchCommand) Queue(args ...any) {
	c.batch.Queue(c.queryString, args...)
}

func (c BatchCommand) Exec() error {
	_, err := c.client.SendBatch(c.ctx, c.batch)
	return err
}
