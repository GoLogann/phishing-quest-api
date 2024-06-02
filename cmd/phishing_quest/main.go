package main

import (
	"database/sql"
	"log"
	"phishing-quest/di"
)

func main() {
	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	container := di.NewContainer(db)

	r := container.Router()
	r.Run(":8082")
}
