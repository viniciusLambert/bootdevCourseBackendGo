package main

import (
	"encoding/json"
	"net/http"

	"github.com/viniciusLambert/bootdevCourseBackendGo/internal/auth"
	"github.com/viniciusLambert/bootdevCourseBackendGo/internal/database"
)

func (cfg *apiConfig) UpdateUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	type requestBody struct {
		Email    string `json:"email"`
		Password string `json:"password"`
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

	decoder := json.NewDecoder(r.Body)
	params := requestBody{}

	if err := decoder.Decode(&params); err != nil {
		respondWithError(w, 500, "error decoding parameters", err)
		return
	}

	hashPassword, err := auth.HashPassword(params.Password)
	if err != nil {
		respondWithError(w, 500, "error hasing password", err)
	}

	user, err := cfg.db.UpdateUser(r.Context(), database.UpdateUserParams{
		ID:             userID,
		Email:          params.Email,
		HashedPassword: hashPassword,
	})
	if err != nil {
		respondWithError(w, 500, "error updating user", err)
		return
	}

	respondWithJSON(w, 200, User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Email:     user.Email,
	})
}
