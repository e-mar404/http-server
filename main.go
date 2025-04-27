package main

import (
	"e-mar404/http-server/internal/handlers"
	"log"
	"net/http"
)

func main() {
	log.Printf("Starting server on http://localhost:8080/app\n")

	cfg := &apiConfig{}
	mux := http.NewServeMux()

	mux.Handle("/app", cfg.middlewareMetricsInc(handlers.App()))
	mux.Handle("GET /app/assets/", cfg.middlewareMetricsInc(handlers.Assets()))
	mux.Handle("GET /api/healthz", handlers.Health())
	mux.Handle("POST /api/validate_chirp", handlers.ValidateChirp())
	mux.Handle("GET /admin/metrics", metricsHandler(cfg))
	mux.Handle("POST /admin/reset", cfg.midlewareMetricsReset(metricsResetHandler()))

	server := http.Server {
		Handler: mux,
		Addr: ":8080",
	}

	log.Printf("Error while serving: %v\n", http.ListenAndServe(server.Addr, server.Handler))
}
