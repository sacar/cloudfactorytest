package covidreport

import (
	"io/ioutil"
	"net/http"
)

func CreateNewGetRequest(url string) (*http.Request, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("x-rapidapi-key", "9efcb0738fmsh921e834f4f341d1p1dbf3djsnd1cc5a1770f9")
	req.Header.Add("x-rapidapi-host", "covid-19-data.p.rapidapi.com")
	return req, nil
}

func Response(req *http.Request) ([]byte, error) {
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
