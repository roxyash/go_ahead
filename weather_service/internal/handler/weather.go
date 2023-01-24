package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/roxyash/go_ahead/weather_service/internal/response"
)

// @Summary Forecast
// @Accept json
// @Produce json
// @Param        date   query     string   false  "Date for which you would like to receive weather data"  Format(RFC3339)
// @Param        city   query     string   false  "The city in which you would like to receive weather data"  Format(string)
// @Success 200 {object} response.ForecastResponse
// @Failure 400,404,500 {object} response.ErrorResponse
// @Failure default {object} response.ErrorResponse
// @Router /forecast [get]
func (h *Handler) forecast(c *gin.Context) {
	dateStr := c.Query("date")
	var date time.Time
	if dateStr == "" {
		date = time.Now()
	} else {
		var err error
		date, err = time.Parse("2006-01-02", dateStr)
		if err != nil {
			h.response.NewErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Invalid date format, use YYYY-MM-DD: %v", err))
			return
		}
	}
	city := c.Query("city")
	if city == "" {
		var err error
		var ip string

		if c.ClientIP() == "::1" {
			ip, err = h.service.Ip.Get()
			if err != nil {
				h.response.NewErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("Error get your global ip: %v", err))
				return
			}
		} else {
			ip = c.ClientIP()
		}

		city, err = h.service.Location.Get(ip, h.cfg.LocationApiKey)
		if err != nil {
			h.response.NewErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Error getting city: %v", err))
			return
		}
	}

	temp, err := h.service.Weather.Get(h.cfg.WeatherApiKey, city, date)
	if err != nil {
		h.response.NewErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("Error getting weather forecast: %v", err))
		return
	}

	c.JSON(http.StatusOK, response.ForecastResponse{
		Temp: temp,
		City: city,
	})
}
