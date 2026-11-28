package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync/atomic"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/primawk/chirpy-go/internal/database"
)

type apiConfig struct {
	db             *database.Queries
	fileserverHits atomic.Int32
	platform       string
}

func main() {
	godotenv.Load()
	dbURL := os.Getenv("DB_URL")
	platform := os.Getenv("PLATFORM")

	db, err := sql.Open("postgres", dbURL)

	if err != nil {
		log.Printf("error reading database: %s\n", err)
		return
	}

	dbQueries := database.New(db)

	const filepathRoot = "."
	const port = "8080"

	apiCfg := &apiConfig{
		db:       dbQueries,
		platform: platform,
	}

	mux := http.NewServeMux()
	mux.Handle("/app/", apiCfg.middlewareMetricsInc(http.FileServer(http.Dir(filepathRoot))))

	mux.HandleFunc("GET /api/healthz", handlerReadiness)
	mux.HandleFunc("GET /api/chirps", apiCfg.handlerFetchChirps)
	mux.HandleFunc("GET /api/chirps/{chirpID}", apiCfg.handlerFetchChirp)
	mux.HandleFunc("POST /api/chirps", apiCfg.handlerCreateChirp)
	mux.HandleFunc("POST /api/users", apiCfg.handlerCreateUser)
	mux.HandleFunc("POST /api/login", apiCfg.handlerLogin)

	mux.HandleFunc("GET /admin/metrics", apiCfg.handlerHit)
	mux.HandleFunc("POST /admin/reset", apiCfg.handlerResetUsers)

	s := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	fmt.Println("Serving on http://localhost:8080")
	log.Fatal(s.ListenAndServe())
}
