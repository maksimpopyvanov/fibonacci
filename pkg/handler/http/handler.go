package handler

import (
	"fibonacci/pkg/repository"
	"net/http"
)

type Handler struct {
	repos *repository.Repository
}

func NewHandler(repos *repository.Repository) *Handler {
	return &Handler{repos: repos}
}

func (h *Handler) Routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/sequence", h.getSequence)
	return mux
}
