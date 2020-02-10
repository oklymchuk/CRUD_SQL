package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	var fconf ConfigData
	var port string

	conff := os.Args[1]

	res := fconf.InitConfig(conff)
	if res {
		port = fmt.Sprintf(":%d", fconf.Port)
	}
	sh := newSQLHuman(&fconf)

	mainRoute := mux.NewRouter()
	apiRoute := mainRoute.PathPrefix("/api/v1").Subrouter()
	apiRoute.HandleFunc("/list", sh.GetAll).Methods(http.MethodGet)
	apiRoute.HandleFunc("/list", sh.Add).Methods(http.MethodPost)
	apiRoute.HandleFunc("/list/{ID}", sh.GetOne).Methods(http.MethodGet)
	apiRoute.HandleFunc("/list/{ID}", sh.UpdateOne).Methods(http.MethodPut)
	apiRoute.HandleFunc("/list/{ID}", sh.DeleteOne).Methods(http.MethodDelete)
	_ = port

	if err := http.ListenAndServe(port, mainRoute); err != nil {
		log.Fatal(err.Error())
	}
}
