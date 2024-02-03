package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/heshan-g/go-api/config"
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

	var firebaseResponseBody struct {
		LocalId string `json:"localId"`
	}

	if err = json.Unmarshal(respBody, &firebaseResponseBody); err != nil {
		msg := fmt.Sprintf("Unexpected error (parsing Firebase response body): %s", err.Error())
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	// ----------------------------------------
	now := time.Now()
	token, err := config.SignJwt(map[string]interface{}{
		"alg": "RS256",
		"iss": os.Getenv("FB_SERVICE_ACCOUNT_EMAIL"),
		"sub": os.Getenv("FB_SERVICE_ACCOUNT_EMAIL"),
		"aud": "https://identitytoolkit.googleapis.com/google.identity.identitytoolkit.v1.IdentityToolkit",
		"exp": now.Add(time.Second * 3600).Unix(),
		"iat": now.Unix(),
		"uid": firebaseResponseBody.LocalId,
	})
	if err != nil {
		msg := fmt.Sprintf("Unexpected error (signing JWT): %s", err.Error())
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	fmt.Println(token)
	// ----------------------------------------

	w.WriteHeader(resp.StatusCode)
	w.Write(respBody)
}
