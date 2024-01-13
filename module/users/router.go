package users

import (
	"github.com/go-chi/chi/v5"
)

func Router(r chi.Router) {
	r.Get("/", getUsersHandler)
}
