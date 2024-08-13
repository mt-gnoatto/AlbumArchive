package db

import (
	"database/sql"
	"fmt"
	"log"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", "postgres://matheus:123456789@postgres/albumtacker?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("You connected to your database.")
}
