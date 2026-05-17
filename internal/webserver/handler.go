package webserver

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func HandleRoutes(r *chi.Mux) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		WriteEntries(w)
	})
}
