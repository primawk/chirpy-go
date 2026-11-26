package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/primawk/chirpy-go/internal/database"
)

func (cfg *apiConfig) handlerCreateUser(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	params := reqEmail{}
	err := decoder.Decode(&params)

	if err != nil {
		content := fmt.Sprintf("Error decoding parameters: %s", err)
		respondWithError(w, http.StatusInternalServerError, content)
		return
	}

	payload := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Email:     params.Email,
	}

	user, err := cfg.db.CreateUser(req.Context(), payload)

	if err != nil {
		content := fmt.Sprintf("Error creating user: %s", err)
		respondWithError(w, http.StatusBadRequest, content)
		return
	}

	respBody := User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Email:     user.Email,
	}

	respondWithJSON(w, http.StatusCreated, respBody)
}

func (cfg *apiConfig) handlerResetUsers(w http.ResponseWriter, req *http.Request) {
	if cfg.platform != "dev" {
		respondWithError(w, http.StatusForbidden, "forbidden endpoint")
		return
	}

	if err := cfg.db.DeleteUsers(req.Context()); err != nil {
		content := fmt.Sprintf("error resetting users: %s\n", err)
		respondWithError(w, http.StatusBadRequest, content)
		return
	}

	w.WriteHeader(http.StatusOK)
}
