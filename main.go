package main

import (
	"log"
	"net/http"

	"github.com/heshan-g/go-api/handlers"
)

func main() {
	http.HandleFunc("/", handlers.RootHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
