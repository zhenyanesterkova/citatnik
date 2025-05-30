package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zhenyanesterkova/citatnik/internal/app/quote"
)

type Repositorie interface {
	Ping() error
	Add(quote *quote.Quote) error
	GetAll() []*quote.Quote
	GetRandom() *quote.Quote
	GetByAuthor(author string) []*quote.Quote
	Delete(id uint64) error
}

type RepositorieHandler struct {
	repo Repositorie
}

func NewRepositorieHandler(
	rep Repositorie,
) *RepositorieHandler {
	return &RepositorieHandler{
		repo: rep,
	}
}

func (rh *RepositorieHandler) InitRouter(router *mux.Router) {
	router.HandleFunc("/ping", rh.Ping).Methods(http.MethodGet)
	router.HandleFunc("/quotes", rh.Get).Methods(http.MethodGet)
	router.HandleFunc("/quotes/random", rh.GetRandom).Methods(http.MethodGet)
	router.HandleFunc("/quotes", rh.Add).Methods(http.MethodPost)
}
