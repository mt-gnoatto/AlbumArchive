package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Init() {
	var err error
	DB, err = sql.Open("postgres", "postgres://matheus:123456789@postgres/albumarchive?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("You connected to your database.")
}
