package webserver

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	chimiddle "github.com/go-chi/chi/v5/middleware"
)

func Start(port int) {
	fmt.Printf("Starting webserver on port %d\n", port)

	InitTemplates()

	var r = chi.NewRouter()
	r.Use(chimiddle.StripSlashes)

	HandleRoutes(r)
	HandleApiRoutes(r)

	err := http.ListenAndServe(":"+strconv.Itoa(port), r)
	if err != nil {
		panic(err)
	}
}
