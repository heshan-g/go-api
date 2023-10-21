package root

import "net/http"

func CreateMux() *http.ServeMux {
	rootMux := http.NewServeMux()

	rootMux.HandleFunc("/health-check", healthCheckHandler)
	rootMux.HandleFunc("/", rootHanlder)

	return rootMux
}
