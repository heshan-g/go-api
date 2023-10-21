package auth

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func signInHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		postHandler(w, r)
	}
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(
			w,
			"Failed to read request body",
			http.StatusBadRequest,
		)
	}

	fmt.Printf("%v\n", string(body))

	respBody := []byte("{ \"userId\": \"123\" }")

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respBody)
}
