package handler

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/zhenyanesterkova/citatnik/internal/app/apperrors"
)

const (
	TextDeleteNotFound = "failed delete: not found quote with this ID"
	TextDeleteBadID    = "failed delete quote: ID should be a number"
)

func (rh *RepositorieHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, TextDeleteBadID, http.StatusBadRequest)
		return
	}

	err = rh.repo.Delete(id)
	if err != nil {
		if errors.Is(err, apperrors.ErrDeleteNotFound) {
			http.Error(w, TextDeleteNotFound, http.StatusBadRequest)
			return
		}
		log.Printf("error delete quote with ID %d from memory: %v", id, err)
		http.Error(w, TextServerError, http.StatusInternalServerError)
		return
	}
}
