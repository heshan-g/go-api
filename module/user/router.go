package user

import "net/http"

func CreateMux() *http.ServeMux {
	userMux := http.NewServeMux()

	userMux.HandleFunc("/user", userHandler)

	return userMux
}
