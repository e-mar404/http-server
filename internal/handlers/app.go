package handlers

import "net/http"

func App() http.Handler {
	return http.StripPrefix("/app", http.FileServer(http.Dir("./")))
}
