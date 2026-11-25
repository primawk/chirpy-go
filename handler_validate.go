package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func handlerValidate(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	params := validateResp{}
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

	respBody := returnCleaned{
		Cleaned_body: cleanChirpBody(params.Body),
	}

	respondWithJSON(w, http.StatusOK, respBody)
}
