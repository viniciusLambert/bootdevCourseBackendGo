package auth

import (
	"fmt"
	"net/http"
	"strings"
)

func GetBearerToken(header http.Header) (string, error) {
	authHeader := header.Get("Authorization")
	if authHeader == "" {
		return "", fmt.Errorf("Authorization Header not found.")
	}

	token := strings.Split(authHeader, " ")
	if len(token) != 2 {
		return "", fmt.Errorf("error: poorly formated token")
	}
	return token[1], nil
}
