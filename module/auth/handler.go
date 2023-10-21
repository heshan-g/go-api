package auth

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type UserSignInCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Response struct {
	UserId      string `json:"userId"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	AccessToken string `json:"accessToken"`
}

func signInHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var userCreds UserSignInCredentials

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(
			w,
			"Failed to parse request body",
			http.StatusInternalServerError,
		)
	}
	if err := json.Unmarshal(body, &userCreds); err != nil {
		http.Error(
			w,
			"Failed to Unmarshal JSON",
			http.StatusInternalServerError,
		)
	}

	fmt.Printf("%+v\n", userCreds)

	resp := Response{
		UserId:      "123",
		Name:        "Alice",
		Email:       "alice@email.com",
		AccessToken: "abc123",
	}

	jsonResp, err := json.Marshal(resp)
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
}
