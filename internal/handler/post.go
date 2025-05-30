package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/zhenyanesterkova/citatnik/internal/app/quote"
)

func (rh *RepositorieHandler) Add(w http.ResponseWriter, r *http.Request) {
	quoteToAdd := quote.Quote{}
	dec := json.NewDecoder(r.Body)

	if err := dec.Decode(&quoteToAdd); err != nil {
		log.Printf("error decode quote to add to memory: %v", err)
		http.Error(w, TextServerError, http.StatusInternalServerError)
		return
	}

	err := rh.repo.Add(&quoteToAdd)
	if err != nil {
		log.Printf("error add quote with ID %d to memory: %v", quoteToAdd.ID, err)
		http.Error(w, TextServerError, http.StatusInternalServerError)
		return
	}
}
