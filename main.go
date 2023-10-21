package main

import (
	"log"
	"net/http"

	"github.com/heshan-g/go-api/module/auth"
	"github.com/heshan-g/go-api/module/root"
)

func main() {
	mainMux := http.NewServeMux()

	mainMux.Handle("/", root.CreateMux())
	mainMux.Handle("/auth/", auth.CreateMux())
	mainMux.Handle("*", http.NotFoundHandler())

	err := http.ListenAndServe(":8080", mainMux)
	if err != nil {
		log.Fatal(err)
	}
}
