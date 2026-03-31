package main

import (
	"sync/atomic"

	"github.com/viniciusLambert/bootdevCourseBackendGo/internal/database"
)

type apiConfig struct {
	fileserverHits atomic.Int32
	platform       string
	db             *database.Queries
}
