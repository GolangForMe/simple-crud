package helpers

import "net/http"

func InitHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}
