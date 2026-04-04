package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/viniciusLambert/bootdevCourseBackendGo/internal/auth"
	"github.com/viniciusLambert/bootdevCourseBackendGo/internal/database"
)

func (cfg *apiConfig) HandleWebhookPolka(w http.ResponseWriter, r *http.Request) {
	type requestBody struct {
		Event string `json:"event"`
		Data  struct {
			UserID string `json:"user_id"`
		} `json:"data"`
	}

	apiKey, err := auth.GetApiKey(r.Header)
	if err != nil {
		respondWithError(w, 401, "error Authentication Token not found", err)
		return
	}

	polkaAPIKEY := os.Getenv("POLKA_SECRET")
	if apiKey != polkaAPIKEY {
		respondWithError(w, 401, "error Authentication Token not found", err)
		return
	}

	decoder := json.NewDecoder(r.Body)
	params := requestBody{}

	if err := decoder.Decode(&params); err != nil {
		respondWithError(w, 500, "error decoding parameters", err)
		return
	}

	if params.Event != "user.upgraded" {
		respondWithJSON(w, 204, nil)
		return
	}

	userID, err := uuid.Parse(params.Data.UserID)
	if err != nil {
		respondWithError(w, 404, "error parsing uuid", err)
		return
	}
	_, err = cfg.db.UpdateUserIsRed(r.Context(), database.UpdateUserIsRedParams{
		ID:          userID,
		IsChirpyRed: true,
	})
	if err != nil {
		respondWithError(w, 404, "error getting chirps from database", err)
		return
	}

	respondWithJSON(w, 204, nil)
}
