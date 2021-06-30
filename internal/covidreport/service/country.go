package service

import (
	"fmt"

	"github.com/sacar/covidreport/internal/covidreport/models"
)

func GetCountryDetails(total models.Total, country models.CountryTotal) models.CountryPercent {
	return models.CountryPercent{
		Country:   country.Country,
		Confirmed: fmt.Sprintf("%.5f %%", percent(country.Confirmed, total.Confirmed)),
		Recovered: fmt.Sprintf("%.5f %%", percent(country.Recovered, total.Recovered)),
		Critical:  fmt.Sprintf("%.5f %%", percent(country.Critical, total.Critical)),
		Deaths:    fmt.Sprintf("%.5f %%", percent(country.Deaths, total.Deaths)),
	}
}

func percent(value int, total int) float64 {
	return float64(value) / float64(total) * 100
}
