package redirect

import (
	"fmt"
	"net/http"

	"github.com/lugosieben/htredirect/config"
)

type Redirection struct {
	Target string
	Method config.Method
}

func check(r *http.Request) (*Redirection, bool) {
	for _, entry := range config.Entries {
		match, err := entry.MatchRequest(r)
		if err != nil {
			fmt.Printf("Error matching request to entry: %s\n", err)
			continue
		}
		if match {
			return &Redirection{
				Target: entry.Target,
				Method: entry.Method,
			}, true
		}
	}
	return &Redirection{}, false
}
