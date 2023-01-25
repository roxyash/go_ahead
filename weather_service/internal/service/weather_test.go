package service

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
)

func TestGetWeatherFree(t *testing.T) {
	if err := godotenv.Load("../../config/.env"); err != nil {
		log.Fatal(err)
	}
	apiKey := os.Getenv("free_weather_apikey")
	// Arrange
	service := new(WeatherService)
	testCases := []struct {
		name        string
		apiKey      string
		city        string
		date        time.Time
		temperature string
		err         string
	}{
		{
			name:        "valid input",
			apiKey:      apiKey,
			city:        "New York",
			date:        time.Date(2023, 1, 25, 0, 0, 0, 0, time.UTC),
			temperature: "3°",
			err:         "",
		},
		{
			name:        "invalid api key",
			apiKey:      "invalid_api_key",
			city:        "New York",
			date:        time.Date(2023, 1, 25, 0, 0, 0, 0, time.UTC),
			temperature: "",
			err:         "API key has been disabled.",
		},
		{
			name:        "no data for specific date",
			apiKey:      apiKey,
			city:        "New York",
			date:        time.Date(2000, 1, 25, 0, 0, 0, 0, time.UTC),
			temperature: "",
			err:         "no data found for specific date",
		},
	}

	// Act and Assert
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			temperature, err := service.GetWeatherFree(tc.apiKey, tc.city, tc.date)
			if err != nil && err.Error() != tc.err {
				t.Errorf("Expected error %s, but got %s", tc.err, err)
			}
			if temperature != tc.temperature {
				t.Errorf("Expected temperature %s, but got %s", tc.temperature, temperature)
			}
		})
	}
}
