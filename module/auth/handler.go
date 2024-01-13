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

	w.WriteHeader(resp.StatusCode)
	w.Write(respBody)
}
