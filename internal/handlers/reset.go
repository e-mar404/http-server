package handlers

import (
	"e-mar404/http-server/internal/api"
	"e-mar404/http-server/internal/respond"
	"fmt"
	"net/http"
)

func ResetHandler(cfg *api.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if cfg.Platform != "dev" {
			respond.Error(w, http.StatusForbidden, "You can only reset in dev")
			return 
		}

		err := cfg.DB.DeleteUsers(r.Context())
		if err != nil {
			respond.Error(w, http.StatusInternalServerError, fmt.Sprintf("could not delete users: %v", err))
			return
		}

		respond.Success(w, r, http.StatusOK, "OK")
	})
}
