package auth

import (
	"net/http"
	"testing"
)

func TestGetBearerToken(t *testing.T) {
	tests := []struct {
		name    string
		header  http.Header
		token   string
		wantErr bool
	}{
		{
			name: "correct token",
			header: http.Header{
				"Authorization": []string{"bearer tokenhash"},
			},
			token:   "tokenhash",
			wantErr: false,
		}, {
			name:    "missing Authorization Header",
			header:  http.Header{},
			token:   "",
			wantErr: true,
		}, {
			name: "malformed bearer token, header with 3 words",
			header: http.Header{
				"Authorization": []string{"bearer functin token"},
			},
			token:   "",
			wantErr: true,
		}, {
			name: "malformed bearer token, header with 1 word",
			header: http.Header{
				"Authorization": []string{"bearerfunctin"},
			},
			token:   "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token, err := GetBearerToken(tt.header)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBearerToken() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && token != tt.token {
				t.Errorf("GetBearerToken() expexts %v, got %v", tt.token, token)
			}
		})
	}
}
