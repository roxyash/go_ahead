package service

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestGet(t *testing.T) {
	if err := godotenv.Load("../../config/.env"); err != nil {
		log.Fatal(err)
	}
	apiKey := os.Getenv("geolocation_apikey")

	service := new(LocationService)
	testCases := []struct {
		name     string
		apiKey   string
		ip       string
		location string
		err      string
	}{
		{
			name:     "valid input",
			apiKey:   apiKey,
			ip:       "8.8.8.8",
			location: "Glenmont",
			err:      "",
		},
		{
			name:     "invalid api key",
			apiKey:   "invalid_api_key",
			ip:       "8.8.8.8",
			location: "",
			err:      "Error getting location data: invalid API key",
		},
		{
			name:     "invalid IP",
			apiKey:   apiKey,
			ip:       "256.256.256.256",
			location: "",
			err:      "Error getting location data: invalid IP address",
		},
	}

	// Act and Assert
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			location, err := service.Get(tc.ip, tc.apiKey)
			if err != nil && err.Error() != tc.err {
				t.Errorf("Expected error %s, but got %s", tc.err, err)
			}
			if location != tc.location {
				t.Errorf("Expected location %s, but got %s", tc.location, location)
			}
		})
	}
}
