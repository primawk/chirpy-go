package main

import (
	"fmt"
	"net/http"
)

func (cfg *apiConfig) handlerHit(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	htmlBody := fmt.Sprintf(`<html>
  <body>
    <h1>Welcome, Chirpy Admin</h1>
    <p>Chirpy has been visited %d times!</p>
  </body>
	</html>`, cfg.fileserverHits.Load())

	_, err := w.Write([]byte(htmlBody))
	if err != nil {
		fmt.Println("error getting the metrics information:", err)
	}
}

// func (cfg *apiConfig) handlerReset(w http.ResponseWriter, req *http.Request) {
// 	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
// 	w.WriteHeader(http.StatusOK)
//
// 	cfg.fileserverHits.Store(0)
// }
