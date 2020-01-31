package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
)

func main() {
	fmt.Println("migrate")
	db, err := sql.Open("mysql", "daming:267552@tcp(127.0.0.1:3306)/address?multiStatements=true")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	driver, _ := mysql.WithInstance(db, &mysql.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://./migrations",
		"mysql",
		driver,
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	err = m.Up()
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}
