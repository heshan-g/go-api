package auth

import (
	"bytes"
	"encoding/json"
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
	url := fmt.Sprintf(
		"https://identitytoolkit.googleapis.com/v1/accounts:signInWithPassword?key=%s",
		os.Getenv("FB_WEB_API_KEY"),
	)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	r.Body.Close()

	signInCreds := r.Context().Value("body").(SignInRequestBody)

	if err := json.Unmarshal(body, &signInCreds); err != nil {
		panic(err)
	}

	jsonStr := []byte(fmt.Sprintf(
		`{"email": "%s", "password": "%s", "returnSecureToken": true}`,
		signInCreds.Email,
		signInCreds.Password,
	))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Response status: ", resp.Status)
	fmt.Println("Response headers: ", resp.Header)
	respBody, _ := io.ReadAll(resp.Body)
	fmt.Println("Response body: ", string(respBody))

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	w.Write(respBody)
}
