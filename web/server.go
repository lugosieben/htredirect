package web

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	chimiddle "github.com/go-chi/chi/v5/middleware"
	"github.com/lugosieben/htredirect/api/redirect"
)

func Start(port int) {
	fmt.Printf("Starting server on port %d\n", port)
	InitTemplates()

	var r = chi.NewRouter()

	r.Use(chimiddle.StripSlashes)
	r.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		done := redirect.TryHandleRequest(w, r)
		if !done {
			Write404(w, r)
			return
		}
	})

	err := http.ListenAndServe(":"+strconv.Itoa(port), r)
	if err != nil {
		panic(err)
	}
}
