package dao

import (
	"damingerdai/address/internal/database"
	"database/sql"
)

var conn *sql.DB

func init() {
	conn = database.GetDataSource()
}
