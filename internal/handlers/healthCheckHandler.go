package handlers

import (
	"fmt"
	"net/http"
)

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprint(w, "hi")
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
