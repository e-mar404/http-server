package main

import (
	"e-mar404/http-server/internal/database"
	"fmt"
	"io"
	"net/http"
	"sync/atomic"
)

type apiConfig struct {
	fileserverHits atomic.Int32
	dbQueries *database.Queries
}

func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("[LOG] running metrics...\n")
		cfg.fileserverHits.Add(1)
		next.ServeHTTP(w, r)
	})
}

func (cfg *apiConfig) midlewareMetricsReset(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("[LOG] reseting metrics...\n")
		cfg.fileserverHits.Swap(0)
		next.ServeHTTP(w, r)
	})
}

func metricsHandler(cfg *apiConfig) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Header.Add("Content-Type", "text/html; charset=utf-8")

		dir := http.Dir("./admin")
		file, err := dir.Open("index.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("<p>error while getting admin metrics page<p>"))
			return
		}
		content, err := io.ReadAll(file)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("<p>error while reading admin metrics page<p>"))
			return
		}

		res := fmt.Sprintf(string(content), cfg.fileserverHits.Load())

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(res))
	})
}

func metricsResetHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Header.Add("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
}
