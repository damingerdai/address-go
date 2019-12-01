package database

import (
	"damingerdai/address/config"
	"database/sql"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

var (
	conn *sql.DB
)

func createDataSourceName(conf *config.DBConfig) (dataSourceName string) {
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

	return
}

func createConnection(conf *config.DBConfig) (*sql.DB, error) {
	driverName := "mysql"
	dataSourceName := createDataSourceName(conf)
	db, err := sql.Open(driverName, dataSourceName)

	if err != nil {
		log.Println(err)
		log.Fatal("Something wrong with created DB connection")
		return nil, errors.Wrap(err, "sql.Connect(driverName, dataSourceName)")
	}

	return db, nil
}

func GetConnection(conf *config.DBConfig) *sql.DB {
	if conn == nil {
		log.Println("No db connection, create a new one")
		conn, _ = createConnection(conf)
		log.Println("Done")
	} else {
		log.Println("db instance is created, use existing one")
	}
	return conn
}
