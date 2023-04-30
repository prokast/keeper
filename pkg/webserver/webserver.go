package webserver

import (
	"net/http"
	"path"
	"text/template"

	"github.com/kooberetis/keeper/pkg/keystore"
	"github.com/kooberetis/keeper/pkg/url"
)

type web struct {
	Url   url.Config
	Local []keystore.CertInfo
}

func Webserver(w http.ResponseWriter, r *http.Request) {
	wb := web{Url: url.Url(), Local: keystore.GetCertInfo()}
	fp := path.Join("web", "index.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w, wb); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	return
}
