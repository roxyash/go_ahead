package service

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/joho/godotenv"

	// weather_service\internal\service\weather_test.go
	mock_service "github.com/roxyash/go_ahead/weather_service/internal/service/mocks"
)

func TestWeatherService_Get(t *testing.T) {
	if err := godotenv.Load("../../config/.env"); err != nil {
		fmt.Println(err)
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock Weather service.
	mockWeather := mock_service.NewMockWeather(ctrl)

	// Set up the expected behavior of the mock service.
	mockWeather.EXPECT().Get(os.Getenv("openweathermap_apikey"), "Moscow", time.Date(2022, time.June, 1, 0, 0, 0, 0, time.UTC)).
		Return("20°", nil)

	// Create a new instance of the WeatherService.
	weatherService := &WeatherService{}


	// Call the Get method and check the result.
	temperature, err := weatherService.Get(os.Getenv("openweathermap_apikey"), "Moscow", time.Date(2022, time.June, 1, 0, 0, 0, 0, time.UTC))
	if err != nil {
		t.Errorf("WeatherService.Get returned an error: %v", err)
	}
	if temperature != "20°" {
		t.Errorf("WeatherService.Get returned the wrong temperature: got %s, want %s", temperature, "20°")
	}
}
