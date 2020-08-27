package main

import (
	"bookmaker/pkg/logger"
	"database/sql"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// Handler ...
type Handler struct {
	logger logger.Logger
	db     *sql.DB
}

// NewHandler ...
func newHandler(newLogger logger.Logger, db *sql.DB) *Handler {
	return &Handler{
		logger: newLogger,
		db:     db,
	}
}

//Routers маршрутизация
func (h *Handler) Routers(r *chi.Mux) *chi.Mux {
	r.Use(middleware.Recoverer)

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/ready", h.ready)
	})

	return r
}

func (h *Handler) ready(w http.ResponseWriter, r *http.Request) {
	if err := h.db.Ping(); err != nil {
		http.Error(w, err.Error(), http.StatusAccepted)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}
