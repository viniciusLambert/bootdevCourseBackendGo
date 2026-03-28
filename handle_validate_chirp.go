package main

import (
	"encoding/json"
	"net/http"
)

func HandleValidateChirp(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	type requestBody struct {
		Body string `json:"body"`
	}
	type responseBody struct {
		Valid bool `json:"valid"`
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

	respondWithJSON(w, 200, responseBody{
		Valid: true,
	})
}
