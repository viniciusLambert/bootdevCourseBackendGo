package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	const filePathRoot = "."
	const port = "8080"

	apiCfg := apiConfig{}
	appHandler := http.FileServer(http.Dir(filePathRoot))

	mux := http.NewServeMux()

	mux.Handle("/app/", http.StripPrefix("/app", apiCfg.middlewareMetricsInc(appHandler)))
	mux.HandleFunc("/healthz", HandleHealth)
	mux.HandleFunc("/metrics", apiCfg.HandleMetrics)
	mux.HandleFunc("/reset", apiCfg.HandleReset)
	srv := http.Server{
		Addr:         ":" + port,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,

		Handler: mux,
	}

	log.Fatal(srv.ListenAndServe())
}
