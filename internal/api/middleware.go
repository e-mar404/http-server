package api

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func (cfg *Config) MiddlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("running metrics...\n")
		cfg.FileserverHits.Add(1)
		next.ServeHTTP(w, r)
	})
}

func (cfg *Config) MidlewareMetricsReset(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("resseting metrics...\n")
		cfg.FileserverHits.Swap(0)
		next.ServeHTTP(w, r)
	})
}

func MetricsHandler(cfg *Config) http.Handler {
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

		res := fmt.Sprintf(string(content), cfg.FileserverHits.Load())

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(res))
	})
}
