package database

import (
	"github.com/damingerdai/address-go/internal/config"
	"testing"
)

func init() {
	conf = &config.DBConfig{
		Host:     "127.0.0.1",
		Port:     "3306",
		User:     "root",
		Password: "267552",
		Dbname:   "mysql",
	}
}

// go test -timeout 30s github.com/damingerdai/address-go/internal/datasource -run TestCreateDataSource -v
func TestCreateDataSource(t *testing.T) {
	if conf == nil {
		t.Error("fail to init conf")
	}

	db, err := CreateDataSource(conf)
	if err != nil {
		t.Error("fail to get db")
	}

	if err = db.Ping(); err != nil {
		t.Error("fail to ping db")
	}
}

func BenchmarkCreateDataSource(b *testing.B) {
	if conf == nil {
		b.Error("fail to init conf")
	}
	for i := 0; i < b.N; i++ {
		db, err := CreateDataSource(conf)
		if err != nil {
			b.Error("fail to get db")
		}
		err = db.Ping()
		if err != nil {
			b.Error("fail to ping db")
		}
		db.Close()
	}
}

//func TestGetDataSource(t *testing.T) {
//	if conf == nil {
//		t.Error("fail to init conf")
//	}
//	db := GetDataSource()
//	defer db.Close()
//	if err := db.Ping(); err != nil {
//		t.Error("fail to ping db")
//	}
//}
//
//func BenchmarkGetDataSource(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		db := GetDataSource()
//		if db == nil {
//			b.Error("fail to get db")
//		}
//		err := db.Ping()
//		if err != nil {
//			b.Error("fail to ping db")
//		}
//		db.Close()
//	}
//}
