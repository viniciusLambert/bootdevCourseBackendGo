package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/viniciusLambert/bootdevCourseBackendGo/internal/auth"
	"github.com/viniciusLambert/bootdevCourseBackendGo/internal/database"
)

func (cfg *apiConfig) HandleLogin(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	type requestBody struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	type responseBody struct {
		User
		Token        string `json:"token"`
		RefreshToken string `json:"refresh_token"`
	}

	accessTokenExpirationTime := time.Hour
	decoder := json.NewDecoder(r.Body)
	params := requestBody{}

	if err := decoder.Decode(&params); err != nil {
		respondWithError(w, 500, "error decoding parameters", err)
		return
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

	tokenJWT, err := auth.MakeJWT(user.ID, cfg.jwtToken, accessTokenExpirationTime)
	if err != nil {
		respondWithError(w, 500, "error creating jwtToken", err)
	}

	refreshToken := auth.MakeRefreshToken()
	_, err = cfg.db.CreateRefreshToken(r.Context(), database.CreateRefreshTokenParams{
		Token:  refreshToken,
		UserID: user.ID,
	})
	if err != nil {
		respondWithError(w, 500, "error registering refreshToken on database", err)
		return
	}

	respondWithJSON(w, 200, responseBody{
		User: User{
			ID:        user.ID,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			Email:     user.Email,
		},
		Token:        tokenJWT,
		RefreshToken: refreshToken,
	})
}
