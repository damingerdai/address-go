package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
)

func main() {
	flag.Parse()
	fmt.Fprintln(os.Stdout, "migrate db script...")
	db, err := sql.Open("mysql", "daming:267552@tcp(127.0.0.1:3306)/address?multiStatements=true")
	if err != nil {
		fmt.Fprintf(os.Stderr, "fail to connect to db: %v", err.Error())
		os.Exit(0)
	}
	if err = db.Ping(); err != nil {
		fmt.Fprintf(os.Stderr, "fail to connect to db: %v", err.Error())
		os.Exit(0)
	}
	driver, _ := mysql.WithInstance(db, &mysql.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://./migrations",
		"mysql",
		driver,
	)
	defer m.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fail to get migrate script: %v", err.Error())
		os.Exit(0)
	}
	switch flag.Arg(0) {
	case "up":
		err := m.Up()
		if err != nil && err.Error() != "no change" {
			fmt.Fprintf(os.Stderr, "fail to run migrate script: %v", err.Error())
			os.Exit(0)
		}
	case "down":
		err := m.Down()
		if err != nil && err.Error() != "no change" {
			fmt.Fprintf(os.Stderr, "fail to run migrate script: %v", err.Error())
			os.Exit(0)
		}
	case "clear":
		err := m.Drop()
		if err != nil && err.Error() != "no change" {
			fmt.Fprintf(os.Stderr, "fail to run migrate script: %v", err.Error())
			os.Exit(0)
		}
	default:
		fmt.Fprintln(os.Stdout, "No run db script")
	}

	fmt.Fprintln(os.Stdout, "run migrate db script")
}
