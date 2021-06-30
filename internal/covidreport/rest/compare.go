package rest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/sacar/covidreport/internal/covidreport"
	"github.com/sacar/covidreport/internal/covidreport/models"
	"github.com/sacar/covidreport/internal/covidreport/service"
)

func CompareCountries(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var countryList models.CountryList
	err := json.Unmarshal(reqBody, &countryList)
	if err != nil {
		log.Println("Cannot read json")
		return
	}
	countryDataList := []models.CountryTotal{}
	for _, country := range countryList.Countries {
		// rate of api call is one request per second
		time.Sleep(2 * time.Second)
		cntryData, err := getCountryData(country)
		if err != nil {
			log.Println(err.Error())
			continue
		}

		countryDataList = append(countryDataList, cntryData)
	}

	countryCompared := service.CompareCountries(countryDataList)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(countryCompared)
}

func getCountryData(country string) (models.CountryTotal, error) {
	url := fmt.Sprintf("https://covid-19-data.p.rapidapi.com/country?name=%s", country)
	req, err := covidreport.CreateNewGetRequest(url)
	if err != nil {
		log.Fatal(err.Error())
		return models.CountryTotal{}, err
	}
	countryRes, err := covidreport.Response(req)

	if err != nil {
		log.Fatal(err.Error())
		return models.CountryTotal{}, err
	}

	var cntryTotal []models.CountryTotal
	json.Unmarshal(countryRes, &cntryTotal)
	if len(cntryTotal) == 0 {
		return models.CountryTotal{}, fmt.Errorf("data for %s not available", country)
	}
	return cntryTotal[0], nil
}
