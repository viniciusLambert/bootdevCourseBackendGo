package main

import (
	"encoding/json"
	"net/http"
	"slices"
	"strings"
)

func HandleValidateChirp(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	type requestBody struct {
		Body string `json:"body"`
	}
	type responseBody struct {
		CleanedBody string `json:"cleaned_body"`
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
	respondWithJSON(w, 200, responseBody{
		CleanedBody: cleanedBody,
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
