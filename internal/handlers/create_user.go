package handlers

import (
	"e-mar404/http-server/internal/api"
	"e-mar404/http-server/internal/respond"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID uuid.UUID `json:"id"`
	Email string `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func CreateUser(cfg *api.Config) http.Handler {
	return http.HandlerFunc (
		func(w http.ResponseWriter, r *http.Request) {
			params := struct {
				Email string `json:"email"`
			}{}
			decoder := json.NewDecoder(r.Body)
			err := decoder.Decode(&params)
			if err != nil {
				respond.Error(w, http.StatusInternalServerError, "error reading response body")
				return 
			}
			user, err := cfg.DB.CreateUser(r.Context(), params.Email)
			if err != nil {
				respond.Error(w, http.StatusInternalServerError, fmt.Sprintf("error creating user: %v\n", err))
				return
			}
			userResponse := User {
				ID: user.ID,
				Email: user.Email,
				CreatedAt: user.CreatedAt,
				UpdatedAt: user.UpdatedAt,
			}
			respond.Success(w, r, http.StatusCreated, userResponse)
		})
}
