package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/viniciusLambert/bootdevCourseBackendGo/internal/auth"
)

func (cfg *apiConfig) HandleLogin(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	type requestBody struct {
		Email            string `json:"email"`
		Password         string `json:"password"`
		ExpiresInSeconds int    `json:"expires_in_seconds"`
	}

	type responseBody struct {
		User
		Token string `json:"token"`
	}

	decoder := json.NewDecoder(r.Body)
	params := requestBody{}

	if err := decoder.Decode(&params); err != nil {
		respondWithError(w, 500, "error decoding parameters", err)
		return
	}

	if params.ExpiresInSeconds == 0 {
		params.ExpiresInSeconds = 60 * 60
	}

	user, err := cfg.db.GetUserByEmail(r.Context(), params.Email)
	if err != nil {
		respondWithError(w, 500, "error getting user from database", err)
		return
	}

	matchPasswords, err := auth.CheckPasswordHash(params.Password, user.HashedPassword)
	if err != nil {
		respondWithError(w, 500, "error checking password", err)
		return
	}

	if !matchPasswords {
		respondWithError(w, 401, "401 Unauthorized", nil)
		return
	}

	tokenJWT, err := auth.MakeJWT(user.ID, cfg.jwtToken, time.Duration(params.ExpiresInSeconds)*time.Second)
	if err != nil {
		respondWithError(w, 500, "error creating jwtToken", err)
	}
	respondWithJSON(w, 200, responseBody{
		User: User{
			ID:        user.ID,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			Email:     user.Email,
		},
		Token: tokenJWT,
	})
}
