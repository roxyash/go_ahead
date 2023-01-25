package service

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type LocationService struct {
}

func NewLocationService() *LocationService {
	return &LocationService{}
}


// Method for get Geolocation by client ip 
func (s *LocationService) Get(ip, apiKey string) (string, error) {
	// URL Location API of api.ipstack
	locationAPI := "http://api.ipstack.com"

	url := fmt.Sprintf("%s/%s?access_key=%s", locationAPI, ip, apiKey)
	res, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("Error getting location data: %v", err)
	}
	defer res.Body.Close()

	var data struct {
		City string `json:"city"`
	}
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		return "", fmt.Errorf("Error decoding location data: %v", err)
	}

	return data.City, nil
}
