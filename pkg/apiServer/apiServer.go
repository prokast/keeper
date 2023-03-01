package apiServer

import (
	"encoding/json"
	"net/http"

	"github.com/infrastate/treasure/pkg/note"
	"github.com/infrastate/treasure/pkg/url"
	"github.com/infrastate/treasure/pkg/utils"
)

// var response = make(map[string]interface{}, 3)
type response struct {
	Message interface{}
}

func (r *response) responseWrapper(v interface{}) string {
	r.Message = v
	resp, _ := json.Marshal(r)
	return string(resp)
}

func ApiServer(w http.ResponseWriter, r *http.Request) {
	resp := response{}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	param := r.URL.Query().Get("q")
	statusCode := 200
	if param == "note" {
		w.Write([]byte(resp.responseWrapper(note.Note())))
	} else if param == "url" {
		w.Write([]byte(resp.responseWrapper(url.Url())))
	} else {
		statusCode = 404
		w.WriteHeader(statusCode)
		w.Write([]byte(resp.responseWrapper("Resource Not Found")))
	}
	utils.InfoLog.Println(r.Method, statusCode, r.Host, r.URL.Path, param)
	return
}
