package dao

import (
	"damingerdai/address/internal/address/database"
	"database/sql"
)

var conn *sql.DB

func init() {
	conn = database.GetDataSource()
}
