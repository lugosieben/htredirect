package redirect

import (
	"fmt"
	"net/http"
	"time"

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
		fmt.Printf("%s - %s - Could not match \"%s %s\" to any entry\n", time.Now().Format(config.TIMEFORMAT), r.RemoteAddr, r.Host, r.URL)
		return false
	}
	ReplaceRedirectionParameters(redirection, r)
	fmt.Printf("%s - %s - Matched \"%s %s\" to entry \"%s %s\"\n", time.Now().Format(config.TIMEFORMAT), r.RemoteAddr, r.Host, r.URL, redirection.Method.String(), redirection.Target)
	ExecuteRedirection(w, redirection)
	return true
}
