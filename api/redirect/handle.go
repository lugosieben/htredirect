package redirect

import (
	"net/http"

	"github.com/lugosieben/htredirect/config"
)

func ExecuteRedirection(w http.ResponseWriter, r *Redirection) {
	w.Header().Set("Location", r.Target)
	switch r.Method {
	case config.MethodPermanent:
		w.WriteHeader(http.StatusMovedPermanently)
	case config.MethodTemporary:
		w.WriteHeader(http.StatusFound)
	}
}

func TryHandleRequest(w http.ResponseWriter, r *http.Request) bool {
	redirection, ok := check(r)
	if !ok {
		return false
	}
	ReplaceRedirectionParameters(redirection, r)
	ExecuteRedirection(w, redirection)
	return true
}
