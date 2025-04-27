package handlers

import "net/http"

func Assets() http.Handler {
	return http.StripPrefix("/app/assets/", http.FileServer(http.Dir("./assets")))	
}
