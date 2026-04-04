package main

import (
	"net/http"

	"github.com/viniciusLambert/bootdevCourseBackendGo/internal/auth"
)

func (cfg *apiConfig) HandleRevokeToken(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

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

	err = cfg.db.RevokeToken(r.Context(), token.Token)
	if err != nil {
		respondWithError(w, 401, "error revoking Token", err)
		return
	}
	respondWithJSON(w, 204, nil)
}
