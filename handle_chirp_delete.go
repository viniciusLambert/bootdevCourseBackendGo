package main

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/viniciusLambert/bootdevCourseBackendGo/internal/auth"
)

func (cfg *apiConfig) HandleDeleteChirp(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	chirpID, err := uuid.Parse(r.PathValue("chirpID"))
	if err != nil {
		respondWithError(w, 404, "error parsing uuid", err)
		return
	}

	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, 401, "error Authentication Token not found", err)
		return
	}

	userID, err := auth.ValidateJWT(token, cfg.jwtToken)
	if err != nil {
		respondWithError(w, 401, "error Authentication Token not valid", err)
		return
	}
	chirp, err := cfg.db.GetChirpByID(r.Context(), chirpID)
	if err != nil {
		respondWithError(w, 404, "error getting chirps from database", err)
		return
	}

	if chirp.UserID != userID {
		respondWithError(w, 403, "operation not allowed", err)
		return
	}

	err = cfg.db.DeleteChirpById(r.Context(), chirp.ID)
	if err != nil {
		respondWithError(w, 500, "error while deleting chirp", err)
		return
	}

	respondWithJSON(w, 204, nil)
}
