package models

type Total struct {
	Confirmed int `json:"confirmed"`
	Recovered int `json:"recovered"`
	Critical  int `json:"critical"`
	Deaths    int `json:"deaths"`
}

type CountryTotal struct {
	Country   string `json:"country"`
	Confirmed int    `json:"confirmed"`
	Recovered int    `json:"recovered"`
	Critical  int    `json:"critical"`
	Deaths    int    `json:"deaths"`
}

type CountryPercent struct {
	Country   string
	Confirmed string
	Recovered string
	Critical  string
	Deaths    string
}

type CountryList struct {
	Countries []string `json:"countries"`
}

type CountryComapre struct {
	HighestCase  string
	LowestCase   string
	HighestDeath string
	LowestDeath  string
}
