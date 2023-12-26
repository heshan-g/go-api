package auth

import "github.com/go-chi/chi/v5"

func Router(r chi.Router) {
	r.Post("/sign-in", signInHandler)
	r.Get("/sign-token", signTokenHandler)
	r.Get("/verify-token", verifyTokenHandler)
}
