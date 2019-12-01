package database

import (
	"damingerdai/address/config"
	"fmt"
	"testing"
)

var conf *config.DBConfig

func init() {
	conf = &config.DBConfig{
		Host:     "127.0.0.1",
		Port:     "3306",
		User:     "root",
		Password: "267552",
		Dbname:   "mysql",
	}
}

func TestGetConnection(t *testing.T) {
	fmt.Println(conf)
	if conf == nil {
		t.Error("fail to init conf")
	}

	conn := GetConnection(conf)
	defer conn.Close()
	if conn == nil {
		t.Error("fail to ge connection from db")
	}
}
