package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sacar/covidreport/internal/covidreport/rest"
)

var (
	router = mux.NewRouter().StrictSlash(true)
)

func main() {

	mapUrls()
	log.Fatal(http.ListenAndServe(":10000", router))
}

func mapUrls() {
	router.HandleFunc("/", rest.Homepage)
	router.HandleFunc("/country/{name}", rest.GetCountryDetails)
	router.HandleFunc("/compare", rest.CompareCountries).Methods("POST")
}
