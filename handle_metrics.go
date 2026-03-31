package main

import (
	"fmt"
	"net/http"
)

func (cfg *apiConfig) GetHits() string {
	return fmt.Sprintf("%d", cfg.fileserverHits.Load())
}

func (cfg *apiConfig) ResetHits() {
	cfg.fileserverHits.And(0)
}

func (cfg *apiConfig) HandleMetrics(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(200)
	_, _ = w.Write([]byte(fmt.Sprintf(`<html>
			<body>
				<h1>Welcome, Chirpy Admin</h1>
				<p>Chirpy has been visited %s times!</p>
			</body>
		</html>`, cfg.GetHits())))
}
