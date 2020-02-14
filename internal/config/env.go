package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	sqlHost string
	sqlPort string
	sqlUser string
	sqlPwd  string
	sqlDB   string
}

var e *Env

func New() *Env {
	if e == nil {
		e = &Env{}
		e.loadConfig()
	}

	return e
}

func (env *Env) loadConfig() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file")
	}
	env.sqlHost = os.Getenv("SQL_HOST")
	env.sqlPort = os.Getenv("SQL_PORT")
	env.sqlUser = os.Getenv("SQL_USER")
	env.sqlPwd = os.Getenv("SQL_PASSWORD")
	env.sqlDB = os.Getenv("SQL_DB")
}

func (env *Env) Host() string {
	return env.sqlHost
}

func (env *Env) Port() string {
	return env.sqlPort
}

func (env *Env) User() string {
	return env.sqlUser
}

func (env *Env) Password() string {
	return env.sqlPwd
}

func (env *Env) Db() string {
	return env.sqlDB
}
