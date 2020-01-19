package database

import (
	"damingerdai/address/internal/config"
	"database/sql"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

var (
	conn           *sql.DB
	conf           *config.DBConfig
	dataSourceName string
)

func createDataSourceName(conf *config.DBConfig) string {
	if dataSourceName == "" {
		var b strings.Builder
		b.WriteString(conf.User)
		b.WriteString(":")
		b.WriteString(conf.Password)
		b.WriteString("@")
		b.WriteString("tcp(")
		b.WriteString(conf.Host)
		b.WriteString(":")
		b.WriteString(conf.Port)
		b.WriteString(")")
		b.WriteString("/")
		b.WriteString(conf.Dbname)
		b.WriteString("?charset=utf8&parseTime=True&loc=Local")
		dataSourceName = b.String()
	}

	return dataSourceName
}

func CreateDataSource(conf *config.DBConfig) (*sql.DB, error) {
	driverName := "mysql"
	dataSourceName := createDataSourceName(conf)
	db, err := sql.Open(driverName, dataSourceName)

	if err != nil {
		log.Println(err)
		log.Fatal("Something wrong with created DB connection")
		return nil, errors.Wrap(err, "sql.Connect(driverName, dataSourceName)")
	}

	err = db.Ping()
	if err != nil {
		log.Println(err)
		log.Fatal("Something wrong with created DB connection")
		return nil, errors.Wrap(err, "sql.Connect(driverName, dataSourceName)")
	}

	return db, nil
}

func GetDataSource() *sql.DB {
	if conf == nil {
		conf = config.GetDBConfig()
	}

	db, err := CreateDataSource(conf)
	if err != nil {
		panic(err)
	}

	return db
}
