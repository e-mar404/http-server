package main

import (
	"database/sql"
	"e-mar404/http-server/internal/api"
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
	platform := os.Getenv("PLATFORM")
	secret := os.Getenv("SECRET")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Error while connecting to database: %v\n", err)
	}
	dbQueries := database.New(db)

	cfg := &api.Config{
		DB:       dbQueries,
		Platform: platform,
		Secret: secret,
	}
	mux := http.NewServeMux()

	mux.Handle("/app", cfg.MiddlewareMetricsInc(handlers.App()))
	mux.Handle("GET /app/assets/", cfg.MiddlewareMetricsInc(handlers.Assets()))
	mux.Handle("GET /api/healthz", handlers.Health())
	mux.Handle("POST /api/chirps", handlers.CreateChirp(cfg))
	mux.Handle("GET /api/chirps", handlers.GetChirps(cfg))
	mux.Handle("GET /api/chirps/{chirpID}", handlers.GetChirpByID(cfg))
	mux.Handle("POST /api/users", handlers.CreateUser(cfg))
	mux.Handle("POST /api/login", handlers.Login(cfg))
	mux.Handle("GET /admin/metrics", api.MetricsHandler(cfg))
	mux.Handle("POST /admin/reset", cfg.MidlewareMetricsReset(handlers.ResetHandler(cfg)))

	server := http.Server{
		Handler: mux,
		Addr:    ":8080",
	}

	log.Printf("Starting server on http://localhost:8080/app\n")
	log.Printf("Error while serving: %v\n", http.ListenAndServe(server.Addr, server.Handler))
}
