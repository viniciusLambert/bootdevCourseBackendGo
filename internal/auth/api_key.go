package auth

import (
	"fmt"
	"net/http"
	"strings"
)

func GetApiKey(header http.Header) (string, error) {
	authHeader := header.Get("Authorization")
	if authHeader == "" {
		return "", fmt.Errorf("authorization Header not found")
	}

	token := strings.Split(authHeader, " ")
	if len(token) != 2 && token[0] != "apiKey" {
		return "", fmt.Errorf("error: poorly formated token")
	}
	return token[1], nil
}
