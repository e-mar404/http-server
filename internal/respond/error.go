package respond

import (
	"log"
	"net/http"
)

func Error(w http.ResponseWriter, code int, msg string) {
	log.Printf("Error: %s\n", msg)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write([]byte(msg))
}
