package gorm

import (
	"gorm.io/gorm"
	"hamburgueria/pkg/starter"
	"strings"
	"sync"
)

var (
	dbs           map[string]*gorm.DB
	initOnce      sync.Once
	isInitialized = false
)

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

	  gorm.Initialize()
	}

Using options:

	gorm.Initialize(
	  gorm.HealthCheckIsCritical(false),
	  //any other options
	)
*/
func Initialize() {
	ensureCreated()
	ensureNotInitialized()

	initClients()
	isInitialized = true
}

/*
GetClient returns a previously registered gorm client

Usage:

	func example() {
	  sqlClient:= gorm.GetClient("example")
	}
*/
func GetClient(name string) *gorm.DB {
	if c, ok := dbs[strings.ToLower(name)]; ok {
		return c
	}
	return nil
}

func initClients() {
	dbs = make(map[string]*gorm.DB)
	databaseConfig := starter.GetDatabasesConfig()

	for dbName, dbConf := range databaseConfig {
		client, err := NewClient(dbConf)
		if err != nil {
			panic(err)
		}

		dbs[strings.ToLower(dbName)] = client
	}
}

func ensureCreated() {
	initOnce.Do(func() {
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
