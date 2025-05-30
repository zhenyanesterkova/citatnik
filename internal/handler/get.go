package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

const (
	TextServerError = "Something went wrong... Server error"
)

func (rh *RepositorieHandler) Ping(w http.ResponseWriter, r *http.Request) {
	err := rh.repo.Ping()

	if err != nil {
		log.Printf("storage is not available: %v", err)
		http.Error(w, TextServerError, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (rh *RepositorieHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	quotes := rh.repo.GetAll()

	enc := json.NewEncoder(w)
	if err := enc.Encode(quotes); err != nil {
		log.Printf("failed encode quotes: %v", err)
		http.Error(w, TextServerError, http.StatusInternalServerError)
		return
	}
}
