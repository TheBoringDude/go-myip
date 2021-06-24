package myip

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"
)

// base api url
const API_URL = "https://my-ip.theboringdude.workers.dev/"

// MyIP data structure
type MyIP struct {
	IP         string  `json:"ip"`
	Geo        MyIPGeo `json:"geo"`
	Asn        int64   `json:"asn"`
	Country    string  `json:"country"`
	City       string  `json:"city"`
	Continent  string  `json:"continent"`
	PostalCode string  `json:"postalCode"`
	Region     string  `json:"region"`
	RegionCode string  `json:"regionCode"`
	Timezone   string  `json:"timezone"`
}

type MyIPGeo struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

// loop wrapper, retries handler, api request
func wrapper(retries, sleepSeconds int) (MyIP, error) {
	response := MyIP{}

	for r := retries; r <= retries; r++ {
		req, err := http.Get(API_URL)
		if err != nil {
			time.Sleep(time.Second * time.Duration(sleepSeconds))
			continue
		}

		defer req.Body.Close()

		body, err := io.ReadAll(req.Body)
		if err != nil {
			return response, errors.New("error reading the response body")
		}

		json.Unmarshal(body, &response)

		return response, nil
	}

	return MyIP{}, errors.New("failed to get the Client's IP, please try again later")
}

// GetMyIP requests the api url and returns an instance of the MyIP which is a parsed response from the api.
func GetMyIP() (MyIP, error) {
	return wrapper(0, 0)
}

// GetMyIPWithRetry requests the api url and returns an instance of the MyIP which is a parsed response from the api.
// If it fails, it will continue with the retries available.
func GetMyIPWithRetry(retries, sleepSeconds int) (MyIP, error) {
	return wrapper(retries, sleepSeconds)
}
