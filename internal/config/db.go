package config

import "os"

type DBConfig struct {
	Host, Port, User, Password, Dbname string
}

var conf *DBConfig

func initDBConfig() {
	a := createDBConfig()
	conf = &a
}

func createDBConfig() DBConfig {
	return DBConfig{Host: os.Getenv("SQL_HOST"), Port: os.Getenv("SQL_PORT"), User: os.Getenv("SQL_USER"), Password: os.Getenv("SQL_PASSWORD"), Dbname: os.Getenv("SQL_DB")}
}

func GetDBConfig() *DBConfig {
	if conf == nil {
		initDBConfig()
	}
	return conf
}
