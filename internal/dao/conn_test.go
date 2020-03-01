package dao

import (
	"damingerdai/address/internal/config"
	"os"
)

var conf *config.DBConfig

func init() {
	conf = &config.DBConfig{
		Host:     "127.0.0.1",
		Port:     "3306",
		User:     "root",
		Password: "267552",
		Dbname:   "address",
	}
	os.Setenv("SQL_HOST", conf.Host)
	os.Setenv("SQL_PORT", conf.Port)
	os.Setenv("SQL_USER", conf.User)
	os.Setenv("SQL_PASSWORD", conf.Password)
	os.Setenv("SQL_DB", conf.Dbname)
}

//func TestGetConnection(t *testing.T) {
//	conn := GetConnection()
//	conn.Ping()
//	defer conn.Close()
//	rows, err := conn.Query("SELECT _id FROM city")
//	if err != nil {
//		t.Error(err)
//	}
//	defer rows.Close()
//}
//
//func BenchmarkGetConnection(b *testing.B) {
//	for i := 1; i < b.N; i++ {
//		conn := GetConnection()
//		conn.Ping()
//		defer conn.Close()
//	}
//}
