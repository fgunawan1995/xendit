package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/pkg/errors"

	"github.com/fgunawan1995/xendit/model"
	"github.com/gorilla/mux"
)

func writeErrorResponse(w http.ResponseWriter, statusCode int, err error) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(model.BuildAPIResponseError(statusCode, errors.Cause(err)))
}

func writeSuccessResponse(w http.ResponseWriter, data interface{}) {
	statusCode := http.StatusOK
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(model.BuildAPIResponseSuccess(statusCode, data))
}

func (h *Handler) GetCharacterByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	characterID := vars["id"]
	if characterID == "" {
		writeErrorResponse(w, http.StatusBadRequest, errors.New("character must be provided"))
		return
	}
	charIDInt, err := strconv.ParseInt(characterID, 10, 64)
	if err != nil {
		writeErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	result, err := h.usecase.GetCharacterByID(charIDInt)
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	writeSuccessResponse(w, result)
}

func (h *Handler) GetAllCharacterIDs(w http.ResponseWriter, r *http.Request) {
	result, err := h.usecase.GetAllCharacterIDs()
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	writeSuccessResponse(w, result)
}
