package handlers

import (
	"e-mar404/http-server/internal/respond"
	"encoding/json"
	"net/http"
	"strings"
)

type rawChirp struct {
	Body string `json:"body"`
}

func ValidateChirp() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		chirp := rawChirp{}
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&chirp)	
		if err != nil {
			respond.Error(w, http.StatusInternalServerError, "Unable to decode chirp")
			return 
		}
		
		valid := len(chirp.Body) <= 140
		if !valid  {
			respond.Error(w, http.StatusBadRequest, "Chirp is too long")
			return
		}

		respond.Success(w, r, http.StatusOK, respond.SuccessResponse {
			CleanedBody: censor(chirp.Body),
		})
	})
}

func censor(str string) string {
	invalidWords := []string {
		"kerfuffle",
		"sharbert",
		"fornax",
	}
	
	for _, word := range invalidWords {
		lower := strings.ToLower(str)
		idx := strings.Index(lower, word)
		if idx == -1 {
			continue
		}
		originalWord := str[idx:idx+len(word)]
		parts := strings.Split(str, originalWord)
		str = strings.Join(parts, "****")
	}

	return str
}
