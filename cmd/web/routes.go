package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() http.Handler {
	mux := chi.NewMux()

	mux.Use(middleware.Recoverer)
	mux.Use(middleware.Timeout(time.Second * 60))

	mux.Route("/api/v1", func(r chi.Router) {
		r.Get("/", app.CheckHealth)
		r.Get("/dog-from-factory", app.CreateDogFromFactory)
		r.Get("/dog-from-abstract-factory", app.CreateDogFromAbstractFactory)
		r.Get("/cat-from-abstract-factory", app.CreateCatFromAbstractFactory)

	})

	return mux
}
