package root

import "net/http"

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		getHealthCheckHandler(w, r)
	} else {
		http.Error(
			w,
			"Request method not allowed",
			http.StatusMethodNotAllowed,
		)
	}
}

func rootHanlder(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		getRootHandler(w, r)
	} else {
		http.Error(
			w,
			"Request method not allowed",
			http.StatusMethodNotAllowed,
		)
	}
}

func getHealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Server is healthy"))
}

func getRootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Go-API"))
}
