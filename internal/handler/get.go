package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/zhenyanesterkova/citatnik/internal/app/quote"
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
}

func (rh *RepositorieHandler) Get(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	author := q.Get("author")

	var quotes []*quote.Quote
	if author == "" {
		quotes = rh.repo.GetAll()
	} else {
		quotes = rh.repo.GetByAuthor(author)
	}

	enc := json.NewEncoder(w)
	if err := enc.Encode(quotes); err != nil {
		log.Printf("failed encode quotes: %v", err)
		http.Error(w, TextServerError, http.StatusInternalServerError)
		return
	}
}

func (rh *RepositorieHandler) GetRandom(w http.ResponseWriter, r *http.Request) {
	quote := rh.repo.GetRandom()

	enc := json.NewEncoder(w)
	if err := enc.Encode(quote); err != nil {
		log.Printf("failed encode random quote: %v", err)
		http.Error(w, TextServerError, http.StatusInternalServerError)
		return
	}
}
