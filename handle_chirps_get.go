package main

import (
	"net/http"
	"sort"

	"github.com/google/uuid"
	"github.com/viniciusLambert/bootdevCourseBackendGo/internal/database"
)

func (cfg *apiConfig) HandleGetChirps(w http.ResponseWriter, r *http.Request) {
	authorId := r.URL.Query().Get("author_id")
	sortParam := r.URL.Query().Get("sort")
	var chirps []database.Chirp
	var err error
	if authorId == "" {
		chirps, err = cfg.db.GetChirps(r.Context())
		if err != nil {
			respondWithError(w, 500, "error getting chirps from database", err)
			return
		}
	} else {
		parsedUserID, err := uuid.Parse(authorId)
		if err != nil {
			respondWithError(w, 500, "error parsing authorID", err)
			return
		}
		chirps, err = cfg.db.GetChirpsByUserID(r.Context(), parsedUserID)
		if err != nil {
			respondWithError(w, 500, "error getting chirps from database", err)
			return
		}
	}

	if sortParam == "desc" {
		sort.Slice(chirps, func(i, j int) bool {
			return chirps[i].CreatedAt.After(chirps[j].CreatedAt)
		})
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
