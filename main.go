package main

import (
	"fmt"
	"log"
	"net/http"
	"sync/atomic"
	"time"
)

type apiConfig struct {
	fileserverHits atomic.Int32
}

func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.fileserverHits.Add(1)
		next.ServeHTTP(w, r)
	})
}

func (cfg *apiConfig) GetHits() string {
	return fmt.Sprintf("Hits: %d\n", cfg.fileserverHits.Load())
}

func (cfg *apiConfig) ResetHits() {
	cfg.fileserverHits.And(0)
}

func main() {
	const filePathRoot = "."
	const port = "8080"

	apiCfg := apiConfig{}
	appHandler := http.FileServer(http.Dir(filePathRoot))

	mux := http.NewServeMux()
	mux.Handle("/app/", http.StripPrefix("/app", apiCfg.middlewareMetricsInc(appHandler)))

	mux.HandleFunc("/healthz", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Add("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(200)
		_, _ = w.Write([]byte("OK"))
	})

	mux.HandleFunc("/metrics", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Add("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(200)
		_, _ = w.Write([]byte(apiCfg.GetHits()))
	})

	mux.HandleFunc("/reset", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Add("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(200)
		apiCfg.ResetHits()
	})
	srv := http.Server{
		Addr:         ":" + port,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,

		Handler: mux,
	}

	log.Fatal(srv.ListenAndServe())
}
