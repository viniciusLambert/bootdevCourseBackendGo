package main

import (
	"encoding/json"
	"net/http"

	"github.com/viniciusLambert/bootdevCourseBackendGo/internal/auth"
)

func (cfg *apiConfig) HandleLogin(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	type requestBody struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

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

	respondWithJSON(w, 200, User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Email:     user.Email,
	})
}
