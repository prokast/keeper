package webserver

import (
	"net/http"
	"path"
	"text/template"

	"github.com/kooberetis/keeper/pkg/url"
)

func Webserver(w http.ResponseWriter, r *http.Request) {
	url := url.Url()
	fp := path.Join("web", "index.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w, url); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	return
}
