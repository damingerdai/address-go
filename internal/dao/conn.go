package dao

import (
	database "damingerdai/address/internal/database"
	"database/sql"
)

var conn *sql.DB

func GetConnection() *sql.DB {
	if conn == nil {
		fetchConnection()
	}

	return conn
}

func fetchConnection() {
	conn = database.GetDataSource()
	err := conn.Ping()
	if err != nil {
		panic(err)
	}
}
