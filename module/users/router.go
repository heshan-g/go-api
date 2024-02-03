package users

import (
	"github.com/go-chi/chi/v5"
	"github.com/heshan-g/go-api/customMiddleware"
)

func Router(r chi.Router) {
	r.Use(customMiddleware.Authenticate)
	r.Get("/", getUsersHandler)
}
