package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/viniciusLambert/bootdevCourseBackendGo/internal/database"
)

func main() {
	const filePathRoot = "."
	const port = "8080"

	godotenv.Load()
	dbURL := os.Getenv("DB_URL")

	apiCfg := apiConfig{}
	apiCfg.db = getDatabaseQueries(dbURL)
	apiCfg.platform = os.Getenv("PLATFORM")

	mux := http.NewServeMux()

	appHandler := http.FileServer(http.Dir(filePathRoot))
	mux.Handle("/app/", http.StripPrefix("/app", apiCfg.middlewareMetricsInc(appHandler)))
	mux.HandleFunc("GET /api/healthz", HandleHealth)

	mux.HandleFunc("POST /api/validate_chirp", HandleValidateChirp)

	mux.HandleFunc("POST /api/users", apiCfg.CreateUser)

	mux.HandleFunc("GET /admin/metrics", apiCfg.HandleMetrics)
	mux.HandleFunc("POST /admin/reset", apiCfg.HandleReset)

	srv := http.Server{
		Addr:         ":" + port,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,

		Handler: mux,
	}

	log.Fatal(srv.ListenAndServe())
}

func getDatabaseQueries(dbURL string) *database.Queries {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("error connecting to database")
	}
	return database.New(db)
}
