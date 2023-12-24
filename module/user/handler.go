package user

import (
	"encoding/json"
	"net/http"
)

type User struct {
	UserId   string `json:"userId"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	IsActive bool   `json:"isActive"`
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getUsers(w, r)
	default:
		http.Error(
			w,
			"Request method not allowed",
			http.StatusMethodNotAllowed,
		)
	}
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	users := []User{
		{
			UserId: "1",
			Name: "Alice",
			Email: "alice@email.com",
			IsActive: true,
		},
		{
			UserId: "2",
			Name: "Bob",
			Email: "bob@email.com",
			IsActive: false,
		},
		{
			UserId: "3",
			Name: "Charlie",
			Email: "charlie@email.com",
			IsActive: true,
		},
	}

	jsonResp, err := json.Marshal(users)
	if err != nil {
		http.Error(
			w,
			"Failed to Marshal JSON",
			http.StatusInternalServerError,
		)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResp)
	return
}
