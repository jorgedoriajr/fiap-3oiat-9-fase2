package sql

import (
	"fmt"
	"hamburgueria/pkg/healthcheck"
	"hamburgueria/pkg/starter"
	"strings"
	"sync"
)

var (
	clients          map[string]Client
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
func Initialize(options ...Opt) {
	ensureCreated()
	ensureNotInitialized()

	for _, op := range options {
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
GetClient returns a previously registered sql client

Usage:

	func example() {
	  sqlClient:= sql.GetClient("example")
	}
*/
func GetClient(name string) Client {
	if c, ok := clients[strings.ToLower(name)]; ok {
		return c
	}

	panic(fmt.Sprintf("no sql client found for %q", name))
}

func initClients() {
	clients = make(map[string]Client)
	databaseConfig := starter.GetDatabasesConfig()

	for dbName, dbConf := range databaseConfig {
		client, err := NewSqlClient(starter.GetAppConfig().Name, dbConf)
		if err != nil {
			fmt.Println(err.Error())
			panic(err)
		}

		clients[strings.ToLower(dbName)] = client
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
