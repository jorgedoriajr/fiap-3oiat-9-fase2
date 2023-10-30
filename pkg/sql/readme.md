`auto generated documentation`

# sql
--
    import "."


## Usage

#### type Batch

```go
type Batch struct {
}
```


#### func  NewBatch

```go
func NewBatch() *Batch
```

#### func (*Batch) Queue

```go
func (b *Batch) Queue(query string, args ...any)
```

#### type BatchCommand

```go
type BatchCommand struct {
	Command
}
```


#### func  NewBatchCommand

```go
func NewBatchCommand(ctx context.Context, client SqlClient, query string) *BatchCommand
```

#### func (BatchCommand) Exec

```go
func (c BatchCommand) Exec() error
```

#### func (BatchCommand) Queue

```go
func (c BatchCommand) Queue(args ...any)
```

#### type Command

```go
type Command struct {
}
```


#### func  NewCommand

```go
func NewCommand(ctx context.Context, client SqlClient, query string, args ...any) *Command
```

#### func (Command) Exec

```go
func (c Command) Exec() error
```

#### type CommandTag

```go
type CommandTag struct {
}
```


#### func (CommandTag) RowsAffected

```go
func (c CommandTag) RowsAffected() int
```

#### func (CommandTag) String

```go
func (c CommandTag) String() string
```

#### type Config

```go
type Config struct {
	Host         string
	Port         int
	DatabaseName string
	User         string
	Password     string
	MaxPoolSize  int
	Schema       string
}
```

Config represents sql database configuration

#### type Query

```go
type Query[T any] struct {
}
```


#### func  NewQuery

```go
func NewQuery[T any](ctx context.Context, client SqlClient, query string, args ...any) *Query[T]
```

#### func (Query[T]) Many

```go
func (q Query[T]) Many() ([]T, error)
```

#### func (Query[T]) One

```go
func (q Query[T]) One() (T, error)
```

#### type Row

```go
type Row interface {
	pgx.Row
}
```


#### type Rows

```go
type Rows interface {
	pgx.Rows
}
```


#### type ScalarCommand

```go
type ScalarCommand[T any] struct {
}
```


#### func  NewScalarCommand

```go
func NewScalarCommand[T any](ctx context.Context, client SqlClient, query string, args ...any) *ScalarCommand[T]
```

#### func (ScalarCommand[T]) Exec

```go
func (c ScalarCommand[T]) Exec() (*T, error)
```

#### type SqlClient

```go
type Client interface {
	Query(ctx context.Context, query string, args ...any) (Rows, error)
	QueryRow(ctx context.Context, query string, args ...any) Row
	Exec(ctx context.Context, query string, args ...any) (CommandTag, error)
	SendBatch(ctx context.Context, batch *Batch) (CommandTag, error)

	Ping(ctx context.Context) error
}
```


#### func  NewSqlClient

```go
func NewSqlClient(conf Config) (SqlClient, error)
```
