package main

import (
	"encoding/json"
	"net/http"
	"slices"
	"strings"

	"github.com/google/uuid"
	"github.com/viniciusLambert/bootdevCourseBackendGo/internal/database"
)

func (cfg *apiConfig) HandleCreateChirpy(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	type requestBody struct {
		Body   string    `json:"body"`
		UserID uuid.UUID `json:"user_id"`
	}

	decoder := json.NewDecoder(r.Body)
	params := requestBody{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 500, "error decoding parameters", err)
		return
	}

	if len(params.Body) > 140 {
		respondWithError(w, 400, "Chirp is too long.", nil)
		return
	}
	cleanedBody := RemoveProfaneWords(params.Body)

	chirpy, err := cfg.db.CreateChirpy(r.Context(), database.CreateChirpyParams{
		Body:   cleanedBody,
		UserID: params.UserID,
	})
	if err != nil {
		respondWithError(w, 500, "error creating chirpy", err)
		return
	}

	respondWithJSON(w, 201, Chirpy{
		ID:        chirpy.ID.String(),
		CreatedAt: chirpy.CreatedAt,
		UpdatedAt: chirpy.UpdatedAt,
		Body:      chirpy.Body,
		UserID:    chirpy.UserID.String(),
	})
}

func RemoveProfaneWords(sentence string) string {
	profaneWords := []string{"kerfuffle", "sharbert", "fornax"}
	words := strings.Split(sentence, " ")

	for i, word := range words {
		lowerCaseWord := strings.ToLower(word)
		if slices.Contains(profaneWords, lowerCaseWord) {
			words[i] = "****"
		}
	}

	return strings.Join(words, " ")
}
