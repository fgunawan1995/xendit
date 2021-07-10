package handler

import (
	"github.com/fgunawan1995/xendit/usecase"
)

type Handler struct {
	usecase usecase.Usecase
}

func New(usecase usecase.Usecase) *Handler {
	return &Handler{
		usecase: usecase,
	}
}
