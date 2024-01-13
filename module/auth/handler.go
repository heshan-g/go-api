package auth

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

type SignInRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func signInHandler(w http.ResponseWriter, r *http.Request) {
	body := r.Context().Value("body").(SignInRequestBody)

	url := fmt.Sprintf(
		"https://identitytoolkit.googleapis.com/v1/accounts:signInWithPassword?key=%s",
		os.Getenv("FB_WEB_API_KEY"),
	)

	jsonStr := []byte(fmt.Sprintf(
		`{"email": "%s", "password": "%s", "returnSecureToken": true}`,
		body.Email,
		body.Password,
	))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		msg := fmt.Sprintf("Unexpected error (creating POST request): %s", err.Error())
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		msg := fmt.Sprintf("Unexpected error (making request): %s", err.Error())
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		msg := fmt.Sprintf("Unexpected error (reading Firebase response body): %s", err.Error())
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(resp.StatusCode)
	w.Write(respBody)
}
