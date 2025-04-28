package main

import (
	"database/sql"
	"e-mar404/http-server/internal/database"
	"e-mar404/http-server/internal/handlers"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	log.Printf("Connecting to database\n")
	godotenv.Load()
	dbURL := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Error while connecting to database: %v\n", err)
	}
	dbQueries := database.New(db)

	cfg := &apiConfig{
		dbQueries: dbQueries,
	}
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

	log.Printf("Starting server on http://localhost:8080/app\n")
	log.Printf("Error while serving: %v\n", http.ListenAndServe(server.Addr, server.Handler))
}
