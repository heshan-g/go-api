package auth

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
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
	var userCreds UserSignInCredentials

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(
			w,
			"Failed to parse request body",
			http.StatusInternalServerError,
		)
		return
	}
	if err := json.Unmarshal(body, &userCreds); err != nil {
		http.Error(
			w,
			"Failed to Unmarshal JSON",
			http.StatusInternalServerError,
		)
		return
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

func signTokenHandler(w http.ResponseWriter, r *http.Request) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "bar",
		"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})
	tokenString, err := token.SignedString([]byte("my_secret_key"))
	if err != nil {
		panic(err)
	}
	w.Write([]byte(tokenString + "\n"))
}

func verifyTokenHandler(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	_, tokenString, found := strings.Cut(authHeader, " ")
	if authHeader == "" || !found {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("No Authorization header\n"))
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte("my_secret_key"), nil
	})
	if err != nil {
		log.Fatal(err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		fmt.Println(claims["foo"], claims["nbf"])
		w.Write([]byte(fmt.Sprintf("%+v\n", claims)))
	} else {
		fmt.Println(err)
	}
}
