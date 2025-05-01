package handlers

import (
	"e-mar404/http-server/internal/api"
	"e-mar404/http-server/internal/database"
	"e-mar404/http-server/internal/models"
	"e-mar404/http-server/internal/respond"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

type rawChirp struct {
	Body   string    `json:"body"`
	UserID uuid.UUID `json:"user_id"`
}

func CreateChirp(cfg *api.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		chirpRequest := rawChirp{}
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&chirpRequest)
		if err != nil {
			respond.Error(w, http.StatusInternalServerError, "Unable to decode chirp")
			return
		}

		valid := len(chirpRequest.Body) <= 140
		if !valid {
			respond.Error(w, http.StatusBadRequest, "Chirp is too long")
			return
		}

		arg := database.CreateChirpParams{
			UserID: uuid.NullUUID{
				UUID:  chirpRequest.UserID,
				Valid: true,
			},
			Body: censor(chirpRequest.Body),
		}
		chirp, err := cfg.DB.CreateChirp(r.Context(), arg)
		chirpResponse := models.Chirp{
			ID:        chirp.ID,
			UserID:    chirp.UserID.UUID,
			Body:      chirp.Body,
			CreatedAt: chirp.CreatedAt,
			UpdatedAt: chirp.UpdatedAt,
		}
		if err != nil {
			respond.Error(w, http.StatusInternalServerError, "Unable to create chirp")
		}
		respond.Success(w, r, http.StatusCreated, chirpResponse)
	})
}

func censor(str string) string {
	invalidWords := []string{
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
		originalWord := str[idx : idx+len(word)]
		parts := strings.Split(str, originalWord)
		str = strings.Join(parts, "****")
	}

	return str
}
