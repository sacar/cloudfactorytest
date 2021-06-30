package service

import (
	"fmt"

	"github.com/sacar/covidreport/internal/covidreport/models"
)

func CompareCountries(countryDataList []models.CountryTotal) models.CountryComapre {

	highestCase, lowestCase, highestDeath, lowestDeath := countryDataList[0], countryDataList[0], countryDataList[0], countryDataList[0]
	for _, country := range countryDataList {
		if country.Confirmed > highestCase.Confirmed {
			highestCase = country
		}

		if country.Confirmed < lowestCase.Confirmed {
			lowestCase = country
		}

		if country.Deaths > highestDeath.Deaths {
			highestDeath = country
		}

		if country.Deaths < lowestDeath.Deaths {
			lowestDeath = country
		}
	}

	return models.CountryComapre{
		HighestCase:  fmt.Sprintf("%s has the highest case", highestCase.Country),
		LowestCase:   fmt.Sprintf("%s has the lowest case", lowestCase.Country),
		HighestDeath: fmt.Sprintf("%s has the highest death", highestDeath.Country),
		LowestDeath:  fmt.Sprintf("%s has the lowest death", lowestDeath.Country),
	}
}
