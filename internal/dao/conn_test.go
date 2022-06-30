package dao

import (
	"github.com/damingerdai/address-go/internal/config"
	"github.com/damingerdai/address-go/internal/database"
	"github.com/jmoiron/sqlx"
	"os"
	"testing"
)

var (
	conf *config.DBConfig
	db   *sqlx.DB
)

func init() {
	conf = &config.DBConfig{
		Host:     "127.0.0.1",
		Port:     "3306",
		User:     "daming",
		Password: "267552",
		Dbname:   "address",
	}
	os.Setenv("SQL_HOST", conf.Host)
	os.Setenv("SQL_PORT", conf.Port)
	os.Setenv("SQL_USER", conf.User)
	os.Setenv("SQL_PASSWORD", conf.Password)
	os.Setenv("SQL_DB", conf.Dbname)
	db, _ = database.CreateDataSource(conf)
	db.Ping()
}

func TestGetConnection(t *testing.T) {
	if db == nil {
		t.Error("no db")
		return
	}
	db.Ping()
	defer db.Close()
	rows, err := db.Query("SELECT _id FROM city")
	if err != nil {
		t.Error(err)
		return
	}
	defer rows.Close()
	t.Log("pass")
}

//
//func BenchmarkGetConnection(b *testing.B) {
//	for i := 1; i < b.N; i++ {
//		conn := GetConnection()
//		conn.Ping()
//		defer conn.Close()
//	}
//}
