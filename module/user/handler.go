package user

import (
	"encoding/json"
	"net/http"

	"github.com/heshan-g/go-api/config"
)

type User struct {
	Id       string `json:"id"`
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
	var db = config.DB

	rows, queryErr := db.Query(`
		SELECT id, name, email, is_active
		FROM users
	`)
	if queryErr != nil {
		panic(queryErr)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		scanErr := rows.Scan(&user.Id, &user.Name, &user.Email, &user.IsActive)
		if scanErr != nil {
			panic(scanErr)
		}
		users = append(users, user)
	}

	resp, err := json.Marshal(users)
	if err != nil {
		http.Error(
			w,
			"Failed to Marshal JSON",
			http.StatusInternalServerError,
		)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
	return
}
