package main

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

func (cfg *apiConfig) HandleGetChirps(w http.ResponseWriter, r *http.Request) {
	chirps, err := cfg.db.GetChirps(r.Context())
	if err != nil {
		respondWithError(w, 500, "error getting chirps from database", err)
		return
	}

	jsonChirps := []Chirpy{}
	for _, chirpy := range chirps {
		jsonChirps = append(jsonChirps, Chirpy{
			ID:        chirpy.ID.String(),
			CreatedAt: chirpy.CreatedAt,
			UpdatedAt: chirpy.UpdatedAt,
			Body:      chirpy.Body,
			UserID:    chirpy.UserID.String(),
		})
	}
	respondWithJSON(w, 200, jsonChirps)
}

func (cfg *apiConfig) HandleGetChirpByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ID: ", r.PathValue("chirpID"))
	chirpID, err := uuid.Parse(r.PathValue("chirpID"))
	if err != nil {
		respondWithError(w, 404, "error parsing uuid", err)
		return
	}

	chirp, err := cfg.db.GetChirpByID(r.Context(), chirpID)
	if err != nil {
		respondWithError(w, 404, "error getting chirps from database", err)
		return
	}
	jsonChirp := Chirpy{
		ID:        chirp.ID.String(),
		CreatedAt: chirp.CreatedAt,
		UpdatedAt: chirp.UpdatedAt,
		Body:      chirp.Body,
		UserID:    chirp.UserID.String(),
	}

	respondWithJSON(w, 200, jsonChirp)
}
