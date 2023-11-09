package sql

import (
	"gorm.io/gorm"
	"hamburgueria/pkg/healthcheck"
	"hamburgueria/pkg/starter"
	"strings"
	"sync"
)

var (
	dbs              map[string]sqlClient
	stdHealthChecker *healthcheck.StandardChecker
	opts             *options
	initOnce         sync.Once
	isInitialized    = false
)

type options struct {
	healthCheckOptions healthCheckOptions
}

type Opt func(*options) error

/*
Initialize initializes sql clients based on configuration

Config example (yaml)

	databases:
	  readWrite:
	    host: localhost
	    port: 5432
	    databaseName: postgres
	    user: postgres
	    password: mMmd38d23
	    maxPoolSize: 5

Usage:

	func main() {
	  //server must be initialized
	  server.Initialize()

	  sql.Initialize()
	}

Using options:

	sql.Initialize(
	  sql.HealthCheckIsCritical(false),
	  //any other options
	)
*/
func Initialize(optionsParam ...Opt) {
	ensureCreated()
	ensureNotInitialized()

	for _, op := range optionsParam {
		err := op(opts)
		if err != nil {
			panic(err)
		}
	}

	initClients()
	initHealthChecker()
	isInitialized = true
}

/*
GetClient returns a previously registered gorm client

Usage:

	func example() {
	  sqlClient:= sql.GetClient("example")
	}
*/
func GetClient(name string) *gorm.DB {
	if c, ok := dbs[strings.ToLower(name)]; ok {
		return c.conn
	}
	return nil
}

func initClients() {
	dbs = make(map[string]sqlClient)
	databaseConfig := starter.GetDatabasesConfig()

	for dbName, dbConf := range databaseConfig {
		client, err := NewClient(dbConf)
		if err != nil {
			panic(err)
		}

		dbs[strings.ToLower(dbName)] = sqlClient{conn: client}
	}
}

func ensureCreated() {
	initOnce.Do(func() {
		stdHealthChecker = &healthcheck.StandardChecker{}
		opts = &options{
			healthCheckOptions: healthCheckOptions{
				name:       "SQL Database",
				isCritical: true,
			},
		}
	})
}

func ensureNotInitialized() {
	if isInitialized {
		panic("sql database already initialized")
	}
}

func ensureInitialized() {
	if !isInitialized {
		panic("sql database must be initialized")
	}
}
