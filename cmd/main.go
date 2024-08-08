package main

import (
	"phishing-quest/adapter/http"
	"phishing-quest/container"
)

func main() {
	cont := container.NewContainer()

	r := http.SetupRouter(cont)

	err := r.Run(":8080")
	if err != nil {
		return
	}
}
