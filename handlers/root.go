package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		getHandler(w)
	} else if r.Method == http.MethodPost {
		postHandler(w, r)
	} else {
		NotFoundHandler(w,r)
	}
}

func getHandler(w http.ResponseWriter) {
	fmt.Fprintf(w, "GET /")
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
	}

	// fmt.Fprintf(w, "POST /")
	fmt.Printf("%v", string(body))
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not found", http.StatusNotFound)
}
