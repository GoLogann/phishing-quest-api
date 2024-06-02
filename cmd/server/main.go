package main

import (
	"database/sql"
	"log"
	"phishing-quest/internal/di"
)

func main() {
	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err = db.Close()
		if err != nil {

		}
	}(db)

	container := di.NewContainer(db)

	r := container.Router()
	err = r.Run(":8082")
	if err != nil {
		return
	}
}
