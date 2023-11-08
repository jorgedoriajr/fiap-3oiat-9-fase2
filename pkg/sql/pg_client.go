package sql

import (
	"context"
	"fmt"
	"hamburgueria/config"
	"net/url"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joomcode/errorx"
	pgxUUID "github.com/vgarvardt/pgx-google-uuid/v5"
)

type Batch struct {
	pgxBatch *pgx.Batch
}

func NewBatch() *Batch {
	return &Batch{pgxBatch: &pgx.Batch{}}
}

type Client interface {
	Query(ctx context.Context, query string, args ...any) (Rows, error)
	QueryRow(ctx context.Context, query string, args ...any) Row
	Exec(ctx context.Context, query string, args ...any) (CommandTag, error)
	SendBatch(ctx context.Context, batch *Batch) (CommandTag, error)

	Ping(ctx context.Context) error
}

type sqlClient struct {
	conn            *pgxpool.Pool
	applicationName string
}

func appendExtraProperty(extra string, property string) string {
	var value string
	if len(extra) > 0 {
		value = "&"
	}

	return fmt.Sprintf("%s%s%s", extra, value, property)
}

func getConnString(conf Config) (string, error) {
	if err := validateConfig(conf); err != nil {
		return "", err
	}

	var extra string
	if conf.MaxPoolSize > 0 {
		extra = appendExtraProperty(extra, fmt.Sprintf("pool_max_conns=%d", conf.MaxPoolSize))
	}

	if len(conf.Schema) > 0 {
		extra = appendExtraProperty(extra, fmt.Sprintf("search_path=%s", conf.Schema))
	}

	if len(conf.QueryMode) > 0 {
		extra = appendExtraProperty(extra, fmt.Sprintf("default_query_exec_mode=%s", conf.QueryMode))
	}

	connUrl := fmt.Sprintf("postgres://%s:%d/%s?%s", conf.Host, conf.Port, conf.DatabaseName, extra)
	databaseURL, err := url.Parse(connUrl)

	if err != nil {
		return "", errorx.Decorate(err, "Failed to parse databaseURL")
	}

	databaseURL.User = url.UserPassword(conf.User, conf.Password)
	return databaseURL.String(), nil
}

func NewSqlClient(applicationName string, conf config.DatabaseConfig) (Client, error) {
	connString, err := getConnString(Config(conf))
	if err != nil {
		return nil, err
	}

	return connect(applicationName, connString)
}

func connect(applicationName, databaseUrl string) (*sqlClient, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*50)
	defer cancel()

	parseConfig, err := pgxpool.ParseConfig(databaseUrl)

	if err != nil {
		return nil, err
	}

	parseConfig.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		pgxUUID.Register(conn.TypeMap())
		return nil
	}

	conn, errConn := pgxpool.NewWithConfig(ctx, parseConfig)

	if errConn != nil {
		return nil, errorx.Decorate(errConn, "Failed to connect to database")
	}

	client := &sqlClient{
		conn:            conn,
		applicationName: applicationName,
	}

	errPing := conn.Ping(ctx)
	if errPing != nil {
		return client, errorx.Decorate(err, "Failed to check database connection")
	}

	return client, err
}
