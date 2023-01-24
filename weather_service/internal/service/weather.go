package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type WeatherService struct {
}

func NewWeatherService() *WeatherService {
	return &WeatherService{}
}

func (s *WeatherService) Get(apiKey, city string, date time.Time) (string, error) {
	// URL Weather API service of api.openweathermap
	weatherAPI := "https://api.openweathermap.org/data/2.5/forecast"

	u, err := url.Parse(weatherAPI)
	if err != nil {
		return "", fmt.Errorf("Error parsing API endpoint: %v", err)
	}

	// Add query params.
	q := u.Query()
	q.Set("q", city)
	q.Set("appid", apiKey)
	q.Set("units", "metric")
	q.Set("dt", strconv.FormatInt(date.Unix(), 10))
	u.RawQuery = q.Encode()

	res, err := http.Get(u.String())
	if err != nil {
		return "", fmt.Errorf("Ошибка получения данных о погоде: %v", err)
	}

	defer res.Body.Close()

	var data struct {
		List []struct {
			Main struct {
				Temp float64 `json:"temp"`
			} `json:"main"`
			Dt int64 `json:"dt"`
		} `json:"list"`
	}

	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		return "", fmt.Errorf("Error decoding weather data: %v", err)
	}

	var temperature float64
	var count int
	for _, item := range data.List {
		time := time.Unix(item.Dt, 0)
		if time.Year() == date.Year() && time.Month() == date.Month() && time.Day() == date.Day() {
			temperature += item.Main.Temp
			count++
		}
	}
	temperature /= float64(count)

	return fmt.Sprintf("%d°", int(temperature)), nil
}
