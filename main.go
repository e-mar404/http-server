package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync/atomic"
)

type apiConfig struct {
	fileserverHits atomic.Int32
}

type chirpPost struct {
	Body string `json:"body"`
}

type errorResponse struct {
	Error string `json:"error"`
}

type successResponse struct {
	Valid bool `json:"valid"`
}

func main() {
	fmt.Printf("Starting server on http://localhost:8080/app\n")

	cfg := &apiConfig{}
	mux := http.NewServeMux()

	mux.Handle("/app", cfg.middlewareMetricsInc(appHandler()))
	mux.Handle("GET /app/assets/", cfg.middlewareMetricsInc(assetsHandler()))
	mux.Handle("GET /api/healthz", healthHandler())
	mux.Handle("POST /api/validate_chirp", validateChirpHandler())
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

func validateChirpHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		chirp := chirpPost{}
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&chirp)	
		if err != nil {
			errorRes := errorResponse {
				Error: "Something went wrong",
			}

			data, err := json.Marshal(errorRes)
			if err != nil {
				log.Printf("Error decoding chirp: %v\n", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(data)
			return
		}
		
		valid := len(chirp.Body) <= 140
		if !valid  {
			errorRes := errorResponse {
				Error: "Chirp is too long",
			}

			data, err := json.Marshal(errorRes)
			if err != nil {
				log.Printf("Error decoding chirp: %v\n", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write(data)
			return
		}

		success := successResponse {
			Valid: true,
		}
		res, err := json.Marshal(success)
		if err != nil {
			log.Printf("Error marshaling response: %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
		}

		r.Header.Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
		
	})
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
