package respond

import (
	"encoding/json"
	"log"
	"net/http"
)

func Success(w http.ResponseWriter, r *http.Request, code int, data any) {
		res, err := json.Marshal(data)
		if err != nil {
			log.Printf("Error marshaling response: %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Server Error"))
			return 
		}
		r.Header.Add("Content-Type", "application/json")
		w.WriteHeader(code)
		w.Write(res)
}
