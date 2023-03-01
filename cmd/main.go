package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/infrastate/treasure/pkg/apiServer"
	"github.com/infrastate/treasure/pkg/utils"
	"github.com/infrastate/treasure/pkg/webserver"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/v1", apiServer.ApiServer)
	r.HandleFunc("/keeper", webserver.Webserver)
	fileServer := http.FileServer(http.Dir("./web/static"))
	r.PathPrefix("/").Handler(http.StripPrefix("/static", fileServer))
	http.Handle("/", r)
	err := http.ListenAndServe(":8080", nil) // устанавливаем порт для прослушивания
	if err != nil {
		utils.ErrorLog.Println(err)
	}
}
