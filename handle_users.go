package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/viniciusLambert/bootdevCourseBackendGo/internal/auth"
	"github.com/viniciusLambert/bootdevCourseBackendGo/internal/database"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Email     string    `json:"email"`
}

func (cfg *apiConfig) CreateUser(w http.ResponseWriter, r *http.Request) {
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

	hashPassword, err := auth.HashPassword(params.Password)
	if err != nil {
		respondWithError(w, 500, "error hasing password", err)
	}

	user, err := cfg.db.CreateUser(r.Context(), database.CreateUserParams{
		Email:          params.Email,
		HashedPassword: hashPassword,
	})
	if err != nil {
		respondWithError(w, 500, "error creating user", err)
		return
	}

	respondWithJSON(w, 201, User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Email:     user.Email,
	})
}
