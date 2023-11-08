package gorm

import (
	"context"
	"fmt"
	"github.com/joomcode/errorx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"hamburgueria/config"
	"hamburgueria/pkg/sql"
	"time"
)

func appendExtraProperty(extra string, property string) string {
	var value string
	if len(extra) > 0 {
		value = " "
	}

	return fmt.Sprintf("%s%s%s", extra, value, property)
}

func getDsn(conf sql.Config) (string, error) {
	if err := sql.ValidateConfig(conf); err != nil {
		return "", err
	}

	var extra string
	if conf.User != "" {
		extra = appendExtraProperty(extra, fmt.Sprintf("user=%s", conf.User))
	}

	if conf.Password != "" {
		extra = appendExtraProperty(extra, fmt.Sprintf("password=%s", conf.Password))
	}

	dsn := fmt.Sprintf(
		`host=%s %s dbname=%s port=%d sslmode=disable TimeZone=America/Sao_Paulo`,
		conf.Host,
		extra,
		conf.DatabaseName,
		conf.Port,
	)

	return dsn, nil
}

func NewClient(conf config.DatabaseConfig) (*gorm.DB, error) {
	dsn, err := getDsn(sql.Config(conf))
	if err != nil {
		return nil, err
	}

	return connect(dsn)
}

func connect(dsn string) (*gorm.DB, error) {
	_, cancel := context.WithTimeout(context.Background(), time.Second*50)
	defer cancel()
	conn, errConn := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
	})
	if errConn != nil {
		return nil, errorx.Decorate(errConn, "Failed to connect to database")
	}
	return conn, nil
}
