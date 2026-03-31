package main

import (
	"fmt"
	"net/http"
)

func (cfg *apiConfig) HandleReset(w http.ResponseWriter, r *http.Request) {
	if cfg.platform != "dev" {
		respondWithError(w, 403, "403 Forbiden", nil)
		return
	}

	fmt.Println("Reseting Metrics")
	cfg.ResetHits()

	fmt.Println("Reseting Users Table")
	err := cfg.db.ResetUsersTable(r.Context())
	if err != nil {
		respondWithError(w, 500, "error reseting usert table", err)
		return
	}

	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(200)
}
