package web

import (
	redditgo "github.com/dimapanfilov/redditGo"
	"github.com/go-chi/chi/v5"
)

func NewHandler(store redditgo.Store) *Handler {
	h := &Handler{
		Mux:   chi.NewMux(),
		store: store,
	}
}

type Handler struct {
	*chi.Mux

	store redditgo.Store
}
