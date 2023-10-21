package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/heshan-g/go-api/config"
	"github.com/heshan-g/go-api/module/auth"
	"github.com/heshan-g/go-api/module/root"
)

func main() {
	config.LoadDotEnv()

	mainMux := http.NewServeMux()

	mainMux.Handle("/", root.CreateMux())
	mainMux.Handle("/auth/", auth.CreateMux())
	mainMux.Handle("*", http.NotFoundHandler())

	portStr := os.Getenv("PORT")
	port, pErr := strconv.Atoi(portStr)
	if pErr != nil {
		log.Fatal("Error parsing PORT from env. ", pErr.Error())
	}

	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("Starting on port %s\n", addr)
	err := http.ListenAndServe(addr, mainMux)
	if err != nil {
		log.Fatal(err)
	}
}
