package api

import (
	"e-mar404/http-server/internal/database"
	"sync/atomic"
)

type Config struct {
	FileserverHits atomic.Int32
	DB *database.Queries
	Platform string
}

