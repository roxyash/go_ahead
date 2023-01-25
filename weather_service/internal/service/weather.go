package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type WeatherService struct {
}

func NewWeatherService() *WeatherService {
	return &WeatherService{}
}

// This method get average temp for today and feature days (max 5)
func forecast(city, apiKey string) (string, error) {
	forecastAPI := "https://api.openweathermap.org/data/2.5/forecast"

	u, err := url.Parse(forecastAPI)
	if err != nil {
		return "", errors.New(fmt.Sprintf("Error parsing forecast API endpoint: %v", err))
	}

	// Add query params.
	q := u.Query()
	q.Set("q", city)
	q.Set("appid", apiKey)
	q.Set("units", "metric")
	u.RawQuery = q.Encode()

	res, err := http.Get(u.String())
	if err != nil {
		return "", errors.New(fmt.Sprintf("Error getting forecast weather data: %v", err))
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
		return "", errors.New(fmt.Sprintf("Error decoding forecast weather data: %v", err))
	}

	now := time.Now()
	var temperature float64
	var count int
	for _, item := range data.List {
		time := time.Unix(item.Dt, 0)
		if time.Year() == now.Year() && time.Month() == now.Month() && time.Day() == now.Day() {
			temperature += item.Main.Temp
			count++
		}
	}

	if count == 0 {
		return "", errors.New("No data found")
	}

	temperature /= float64(count)

	return fmt.Sprintf("%v°", int(temperature)), nil
}

// This method get data for current temp
func weather(city, apiKey string) (string, error) {
	weatherAPI := "https://api.openweathermap.org/data/2.5/weather"

	u, err := url.Parse(weatherAPI)
	if err != nil {
		return "", errors.New(fmt.Sprintf("Error parsing current API endpoint: %v", err))
	}

	q := u.Query()
	q.Set("q", city)
	q.Set("appid", apiKey)
	q.Set("units", "metric")
	u.RawQuery = q.Encode()

	res, err := http.Get(u.String())
	if err != nil {
		return "", errors.New(fmt.Sprintf("Error getting current weather data: %v", err))
	}

	defer res.Body.Close()

	var currentData struct {
		Main struct {
			Temp float64 `json:"temp"`
		} `json:"main"`
	}

	if err := json.NewDecoder(res.Body).Decode(&currentData); err != nil {
		return "", errors.New(fmt.Sprintf("Error decoding current weather data: %v", err))
	}
	return fmt.Sprintf("%v°", int(currentData.Main.Temp)), nil
}

// Struct for unmarshalling bytes for Paid API service.
type WeatherData struct {
	List []struct {
		Main struct {
			Temp float64 `json:"temp"`
		} `json:"main"`
	} `json:"list"`
}

// This method get history data, for paid Weather API service
func (s *WeatherService) GetOpenWeatherPaid(apiKey, city string, date time.Time) (string, error) {
	// Get start and end times for the date
	start := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
	end := time.Date(date.Year(), date.Month(), date.Day(), 23, 59, 59, 0, date.Location())

	// Format start and end times as yyyy-mm-dd hh:mm:ss
	startStr := start.Format("2006-01-02 15:04:05")
	endStr := end.Format("2006-01-02 15:04:05")

	// Call the API
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/history/city?q=%s&type=hour&start=%s&end=%s&appid=%s", city, startStr, endStr, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read and parse the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var data WeatherData
	err = json.Unmarshal(body, &data)
	if err != nil {
		return "", err
	}
	// Calculate average temperature
	var tempSum float64
	for _, item := range data.List {
		tempSum += item.Main.Temp
	}
	avgTemp := tempSum / float64(len(data.List))

	return fmt.Sprintf("%v°", int(avgTemp)), nil
}

// This method, call two (forecast, weather). First he call forecast method, for check avg temp for specific date,
// if data is null, he call weather method, for get current temp
func (s *WeatherService) GetOpenWeatherFree(apiKey, city string, date time.Time) (string, error) {
	// URL Weather API service of api.openweathermap
	var temperature string
	temperature, err := forecast(city, apiKey)
	if err != nil {
		temperature, err := weather(city, apiKey)
		if err != nil {
			return "", errors.New(fmt.Sprintf("Error in weather method: %v", err))
		}
		return temperature, nil
	} else {
		return temperature, nil
	}
}

// Struct for unmarshalling bytes from Free Weather API Service
type WeatherDataFree struct {
	Forecast struct {
		Forecastday []struct {
			Date string `json:"date"`
			Day  struct {
				Avghumidity float64 `json:"avghumidity"`
				Maxtemp_c   float64 `json:"maxtemp_c"`
				Mintemp_c   float64 `json:"mintemp_c"`
				Avgtemp_c   float64 `json:"avgtemp_c"`
			} `json:"day"`
		} `json:"forecastday"`
	} `json:"forecast"`
}

type ErrorWeatherService struct {
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

// This method return avg temp for specific date/city, from FREE Weather API Service
func (s *WeatherService) GetWeatherFree(apiKey, city string, date time.Time) (string, error) {
	dateStr := date.Format("2006-01-02")

	url := fmt.Sprintf("http://api.weatherapi.com/v1/history.json?key=%v&q=%s&dt=%s", apiKey, city, dateStr)

	url = strings.ReplaceAll(url, " ", "%20")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err

	}
	weatherData := WeatherDataFree{}
	errorMessage := ErrorWeatherService{}
	json.Unmarshal(body, &errorMessage)
	if errorMessage.Error.Message == "" {
		json.Unmarshal(body, &weatherData)
	} else {
		if errorMessage.Error.Code == 1007 {
			return "", errors.New("no data found for specific date")
		}
		return "", errors.New(errorMessage.Error.Message)
	}
	return fmt.Sprintf("%v°", int(weatherData.Forecast.Forecastday[0].Day.Avgtemp_c)), nil
}
