package auth

import "net/http"

func CreateMux() *http.ServeMux {
	authMux := http.NewServeMux()

	authMux.HandleFunc("/auth/sign-in", signInHandler)

	return authMux
}
