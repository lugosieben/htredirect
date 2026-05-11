package webserver

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/lugosieben/htredirect/config"
)

var templateEntries *template.Template

func InitTemplates() {
	templateEntries = template.Must(template.ParseFiles("web/templates/entries.html"))
}

func WriteEntries(w http.ResponseWriter) {
	err := templateEntries.Execute(w, config.Entries)
	if err != nil {
		fmt.Println("Error executing Entries template:", err)
	}
}
