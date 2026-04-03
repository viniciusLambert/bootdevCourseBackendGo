package auth

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestCheckJWT(t *testing.T) {
	tokenSecret1 := "secret1"
	user1uuid := uuid.New()
	tokenSecret2 := "secret2"
	user2uuid := uuid.New()

	jwt1, _ := MakeJWT(user1uuid, tokenSecret1, 10*time.Second)
	jwt2, _ := MakeJWT(user2uuid, tokenSecret2, 10*time.Second)
	expiredJWT, _ := MakeJWT(user1uuid, tokenSecret1, -1*time.Second)

	tests := []struct {
		name        string
		tokenSecret string
		jwt         string
		userID      uuid.UUID
		wantErr     bool
	}{
		{
			name:        "Correct JWT",
			tokenSecret: tokenSecret1,
			jwt:         jwt1,
			userID:      user1uuid,
			wantErr:     false,
		},
		{
			name:        "Incorrect JWT",
			tokenSecret: "wrongSecret",
			jwt:         jwt1,
			userID:      user1uuid,
			wantErr:     true,
		},
		{
			name:        "Secret for a different JWT",
			tokenSecret: tokenSecret1,
			jwt:         jwt2,
			userID:      user1uuid,
			wantErr:     true,
		},
		{
			name:        "Empty secret",
			tokenSecret: "",
			jwt:         jwt1,
			userID:      user1uuid,
			wantErr:     true,
		},
		{
			name:        "Invalid JWT",
			tokenSecret: tokenSecret1,
			jwt:         "invalidJWT",
			userID:      user1uuid,
			wantErr:     true,
		},
		{
			name:        "Expired JWT",
			tokenSecret: tokenSecret1,
			jwt:         expiredJWT,
			userID:      user1uuid,
			wantErr:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotUserID, err := ValidateJWT(tt.jwt, tt.tokenSecret)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateJWT() error = %v, wantErr = %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && gotUserID != tt.userID {
				t.Errorf("validateJWT() expects %v, got %v", tt.userID, gotUserID)
			}
		})
	}
}
