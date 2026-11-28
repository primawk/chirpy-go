package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/primawk/chirpy-go/internal/database"
)

func (cfg *apiConfig) handlerCreateChirp(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	params := reqCreateChirp{}
	err := decoder.Decode(&params)

	if err != nil {
		content := fmt.Sprintf("Error decoding parameters: %s", err)
		respondWithError(w, http.StatusInternalServerError, content)
		return
	}
	if len(params.Body) > 140 {
		respondWithError(w, http.StatusBadRequest, "Chirp is too long")
		return
	}

	payload := database.CreateChirpParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Body:      cleanChirpBody(params.Body),
		UserID:    params.UserId,
	}

	chirp, err := cfg.db.CreateChirp(req.Context(), payload)

	if err != nil {
		content := fmt.Sprintf("Error creating chirp: %s", err)
		respondWithError(w, http.StatusBadRequest, content)
		return
	}

	chirpCreated := Chirp{
		ID:        chirp.ID,
		CreatedAt: chirp.CreatedAt,
		UpdatedAt: chirp.UpdatedAt,
		Body:      chirp.Body,
		UserID:    chirp.UserID,
	}

	respondWithJSON(w, http.StatusCreated, chirpCreated)
}

func (cfg *apiConfig) handlerFetchChirps(w http.ResponseWriter, req *http.Request) {
	chirps, err := cfg.db.GetChirps(req.Context())

	if err != nil {
		content := fmt.Sprintf("Error fetching chirps: %s", err)
		respondWithError(w, http.StatusInternalServerError, content)
		return
	}

	var result []Chirp

	for _, chirp := range chirps {
		chirpJson := Chirp{
			ID:        chirp.ID,
			CreatedAt: chirp.CreatedAt,
			UpdatedAt: chirp.UpdatedAt,
			Body:      chirp.Body,
			UserID:    chirp.UserID,
		}
		result = append(result, chirpJson)
	}

	respondWithJSON(w, http.StatusOK, result)
}

func (cfg *apiConfig) handlerFetchChirp(w http.ResponseWriter, req *http.Request) {
	chirpId, err := uuid.Parse(req.PathValue("chirpID"))

	if err != nil {
		content := fmt.Sprintf("Error getting the chirp id: %s", err)
		respondWithError(w, http.StatusBadRequest, content)
		return
	}

	chirp, err := cfg.db.GetChirp(req.Context(), chirpId)

	if err != nil {
		content := fmt.Sprintf("Error fetching chirp: %s", err)
		respondWithError(w, http.StatusNotFound, content)
		return
	}

	result := Chirp{
		ID:        chirp.ID,
		CreatedAt: chirp.CreatedAt,
		UpdatedAt: chirp.UpdatedAt,
		Body:      chirp.Body,
		UserID:    chirp.UserID,
	}

	respondWithJSON(w, http.StatusOK, result)
}
