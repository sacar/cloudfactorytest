package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/sacar/covidreport/internal/covidreport"
	"github.com/sacar/covidreport/internal/covidreport/models"
	"github.com/sacar/covidreport/internal/covidreport/service"
)

func GetCountryDetails(w http.ResponseWriter, r *http.Request) {

	req, err := covidreport.CreateNewGetRequest("https://covid-19-data.p.rapidapi.com/totals")
	if err != nil {
		log.Println("Error creating Total request")
		log.Fatal(err.Error())
		return
	}
	res, err := covidreport.Response(req)
	if err != nil {
		log.Println("Error creating Total request")
		log.Fatal(err.Error())
		return
	}
	var total []models.Total
	json.Unmarshal(res, &total)

	if len(total) == 0 {
		log.Println("Total record api did not return the data")
		return
	}

	// rate of api call is one request per second
	time.Sleep(2 * time.Second)

	vars := mux.Vars(r)
	url := fmt.Sprintf("https://covid-19-data.p.rapidapi.com/country?name=%s", vars["name"])
	req, err = covidreport.CreateNewGetRequest(url)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	countryRes, err := covidreport.Response(req)

	if err != nil {
		log.Fatal(err.Error())
		return
	}

	var cntryTotal []models.CountryTotal
	json.Unmarshal(countryRes, &cntryTotal)
	if len(cntryTotal) == 0 {
		log.Println("Country Total record api did not return the data")
		return
	}

	response := service.GetCountryDetails(total[0], cntryTotal[0])

	json.NewEncoder(w).Encode(response)
}
