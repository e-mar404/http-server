package handlers

import (
	"e-mar404/http-server/internal/api"
	"e-mar404/http-server/internal/auth"
	"e-mar404/http-server/internal/database"
	"e-mar404/http-server/internal/models"
	"e-mar404/http-server/internal/respond"
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateUser(cfg *api.Config) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			params := struct {
				Email    string `json:"email"`
				Password string `json:"password"`
			}{}
			decoder := json.NewDecoder(r.Body)
			err := decoder.Decode(&params)
			if err != nil {
				respond.Error(w, http.StatusInternalServerError, "error reading response body")
				return
			}
			hash, err := auth.HashPassword(params.Password)
			if err != nil {
				respond.Error(w, http.StatusInternalServerError, err.Error())
				return
			}
			arg := database.CreateUserParams{
				Email:          params.Email,
				HashedPassword: hash,
			}
			user, err := cfg.DB.CreateUser(r.Context(), arg)
			if err != nil {
				respond.Error(w, http.StatusInternalServerError, fmt.Sprintf("error creating user: %v\n", err))
				return
			}
			userResponse := models.User{
				ID:        user.ID,
				Email:     user.Email,
				CreatedAt: user.CreatedAt,
				UpdatedAt: user.UpdatedAt,
			}
			respond.Success(w, r, http.StatusCreated, userResponse)
		})
}
