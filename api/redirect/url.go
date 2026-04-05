package redirect

import (
	"net/http"
	"strings"
)

func ReplaceParameters(url string, r *http.Request) string {
	path := strings.Trim(r.URL.Path, "/")
	return strings.Replace(url, "{path}", path, -1)
}

func ReplaceRedirectionParameters(redirection *Redirection, r *http.Request) {
	redirection.Target = ReplaceParameters(redirection.Target, r)
}
