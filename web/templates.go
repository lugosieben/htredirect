package web

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/lugosieben/htredirect/config"
)

type TemplateData struct {
	Version string
	Host    string
}

var template404 *template.Template

func InitTemplates() {
	template404 = template.Must(template.ParseFiles("web/templates/404.html"))
}

func Write404(w http.ResponseWriter, r *http.Request) {
	err := template404.Execute(w, struct {
		Version string
		Host    string
	}{
		Version: config.VERSION,
		Host:    r.Host,
	})
	if err != nil {
		fmt.Println("Error executing 404 template:", err)
	}
}
