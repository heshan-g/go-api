package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/heshan-g/go-api/config"
	"github.com/heshan-g/go-api/module/auth"
	"github.com/heshan-g/go-api/module/users"
)

func main() {
	config.LoadDotEnv()
	config.ConnectToDb()
	defer config.DB.Close()

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/health-check", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Server is healthy\n"))
	})
	r.Route("/auth", auth.Router)
	r.Route("/users", users.Router)

	portStr := os.Getenv("PORT")
	port, pErr := strconv.Atoi(portStr)
	if pErr != nil {
		log.Fatal("Error parsing PORT from env. ", pErr.Error())
	}

	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("Listening on port %s\n--\n", addr)
	err := http.ListenAndServe(addr, r)
	if err != nil {
		log.Fatal(err)
	}
}
