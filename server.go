package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	addr := flag.String("a", ":8000", "address of app")
	flag.Parse()
	sh := newSQLHuman()

	mainRoute := mux.NewRouter()
	apiRoute := mainRoute.PathPrefix("/api/v1").Subrouter()
	apiRoute.HandleFunc("/list", sh.GetAll).Methods(http.MethodGet)
	apiRoute.HandleFunc("/list", sh.Add).Methods(http.MethodPost)
	apiRoute.HandleFunc("/list/{ID}", sh.GetOne).Methods(http.MethodGet)
	apiRoute.HandleFunc("/list/{ID}", sh.UpdateOne).Methods(http.MethodPut)
	apiRoute.HandleFunc("/list/{ID}", sh.DeleteOne).Methods(http.MethodDelete)

	if err := http.ListenAndServe(*addr, mainRoute); err != nil {
		log.Fatal(err.Error())
	}
}
