package handlers

import (
	"e-mar404/http-server/internal/api"
	"e-mar404/http-server/internal/models"
	"e-mar404/http-server/internal/respond"
	"log"
	"net/http"
)

func GetChirps(cfg *api.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		chirps, err := cfg.DB.GetAllChirps(r.Context())
		if err != nil {
			log.Fatalln(err)
			respond.Error(w, http.StatusInternalServerError, "Could not retrieve chirps")
			return
		}
		
		chirpsResponse := make(models.ChirpList, len(chirps))
		for i, chirp := range chirps {
			chirpsResponse[i] = models.Chirp{
				ID: chirp.ID,
				UserID: chirp.UserID.UUID,
				Body: chirp.Body, 
				CreatedAt: chirp.CreatedAt, 
				UpdatedAt: chirp.UpdatedAt, 
			}
		}
		respond.Success(w, r, http.StatusOK, chirpsResponse)
	})
}
