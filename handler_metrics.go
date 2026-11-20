package main

import (
	"fmt"
	"net/http"
)

func (cfg *apiConfig) handlerHit(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	body := fmt.Sprintf("Hits: %d", (int(cfg.fileserverHits.Load())))

	_, err := w.Write([]byte(body))
	if err != nil {
		fmt.Println("error getting the metrics information:", err)
	}
}

func (cfg *apiConfig) handlerReset(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	cfg.fileserverHits.Store(0)
}
