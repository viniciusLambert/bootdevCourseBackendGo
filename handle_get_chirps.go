package main

import "net/http"

func (cfg *apiConfig) HandleGetChirps(w http.ResponseWriter, r *http.Request) {
	chirps, err := cfg.db.GetChirps(r.Context())
	if err != nil {
		respondWithError(w, 500, "error getting chirps from database", err)
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
