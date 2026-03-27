package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	const filePathRoot = "."
	const port = "8080"

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(filePathRoot)))

	srv := http.Server{
		Addr:         ":" + port,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,

		Handler: mux,
	}

	log.Fatal(srv.ListenAndServe())
}
