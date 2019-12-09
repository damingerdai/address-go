package dao

import (
	"damingerdai/address/database"
	"database/sql"
)

var conn *sql.DB

func init() {
	conn = database.GetDataSource()
}
