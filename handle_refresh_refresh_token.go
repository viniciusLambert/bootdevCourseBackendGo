package main

import (
	"net/http"
	"time"

	"github.com/viniciusLambert/bootdevCourseBackendGo/internal/auth"
)

func (cfg *apiConfig) HandleRefresh(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	type responseBody struct {
		Token string `json:"token"`
	}

	headerToken, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, 401, "error Authentication Token not found", err)
		return
	}

	token, err := cfg.db.GetRefreshTokenByToken(r.Context(), headerToken)
	if err != nil {
		respondWithError(w, 401, "Token not found", err)
		return
	}

	if token.ExpiredAt.Time.Before(time.Now()) {
		respondWithError(w, 401, "Token expired", nil)
		return
	}

	if token.RevokedAt.Valid {
		respondWithError(w, 401, "Token was revoked", nil)
		return
	}

	jwtToken, err := auth.MakeJWT(token.UserID, cfg.jwtToken, time.Hour)
	if err != nil {
		respondWithError(w, 500, "error creating jwtToken", err)
		return
	}

	respondWithJSON(w, 200, responseBody{
		Token: jwtToken,
	})
}
