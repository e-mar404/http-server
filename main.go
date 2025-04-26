package main

import (
	"fmt"
	"io"
	"net/http"
	"sync/atomic"
)

type apiConfig struct {
	fileserverHits atomic.Int32
}

func main() {
	fmt.Printf("Starting server on http://localhost:8080/app\n")

	cfg := &apiConfig{}
	mux := http.NewServeMux()

	mux.Handle("/app", cfg.middlewareMetricsInc(appHandler()))
	mux.Handle("GET /app/assets/", cfg.middlewareMetricsInc(assetsHandler()))
	mux.Handle("GET /api/healthz", healthHandler())
	mux.Handle("GET /admin/metrics", metricsHandler(cfg))
	mux.Handle("POST /admin/reset", cfg.midlewareMetricsReset(metricsResetHandler()))

	server := http.Server {
		Handler: mux,
		Addr: ":8080",
	}

	fmt.Printf("Error while serving: %v\n", http.ListenAndServe(server.Addr, server.Handler))
}

func appHandler() http.Handler {
	return http.StripPrefix("/app", http.FileServer(http.Dir("./")))	
}

func assetsHandler() http.Handler {
	return http.StripPrefix("/app/assets/", http.FileServer(http.Dir("./assets")))	
}

func healthHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Header.Add("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
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
		}
		content, err := io.ReadAll(file)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("<p>error while reading admin metrics page<p>"))
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
