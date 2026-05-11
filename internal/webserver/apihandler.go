package webserver

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func HandleApiRoutes(r *chi.Mux) {
	r.Route("/api", func(r chi.Router) {

		r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})

	})
}
