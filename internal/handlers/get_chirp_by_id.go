package handlers

import (
	"e-mar404/http-server/internal/api"
	"e-mar404/http-server/internal/models"
	"e-mar404/http-server/internal/respond"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

func GetChirpByID(cfg *api.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		chirpID, err := uuid.Parse(r.PathValue("chirpID"))
		if err != nil {
			respond.Error(w, http.StatusInternalServerError, "That is not a valid uuid")
			return
		}

		chirp, err := cfg.DB.GetChirpByID(r.Context(), chirpID)
		if !chirp.UserID.Valid {
			msg := fmt.Sprintf("Chirp with id %v not found", chirpID)
			respond.Error(w, http.StatusNotFound, msg)
			return
		}

		chirpResponse := models.Chirp {
			ID: chirp.ID,
			UserID: chirp.UserID.UUID,
			Body: chirp.Body, 
			CreatedAt: chirp.CreatedAt, 
			UpdatedAt: chirp.UpdatedAt, 
		}
		respond.Success(w, r, http.StatusOK, chirpResponse)
	})
}
