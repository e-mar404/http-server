package handlers

import (
	"e-mar404/http-server/internal/api"
	"e-mar404/http-server/internal/auth"
	"e-mar404/http-server/internal/models"
	"e-mar404/http-server/internal/respond"
	"encoding/json"
	"net/http"
)

func Login(cfg *api.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
		user, err := cfg.DB.GetUserByEmail(r.Context(), params.Email)
		if err != nil {
			respond.Error(w, http.StatusUnauthorized, "Did not find a user with that email and password")
			return
		}

		err = auth.CheckPasswordHash(user.HashedPassword, params.Password)
		if err != nil {
			respond.Error(w, http.StatusUnauthorized, err.Error())
			return
		}

		userResponse := models.User{
			ID:        user.ID,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		}
		respond.Success(w, r, http.StatusOK, userResponse)
	})
}
