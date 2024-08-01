package main

import (
	"phishing-quest/adapter/http"
	"phishing-quest/postgres"
)

func main() {
	postgres.InitDB()

	r := http.SetupRouter()

	err := r.Run(":8080")
	if err != nil {
		panic("failed to start server")
	}
}
