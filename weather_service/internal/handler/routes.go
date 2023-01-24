package handler

// Задача
// Сделать API сервис который по запросу ([GET] /forecast) возвращает прогноз погоды в виде json.
// Сервис должен принимать query параметры:
// - date — дата дня прогноза (по умолчанию текущий день);
// - city — город (по умолчанию сервис должен вычислять текущее местоположение клиента).

// Ответ сервиса
// В ответе должна быть средняя температура за день в градусах и местоположение.
// Для получения погоды и геолокации можно использовать любые API, например:
// - https://openweathermap.org/
// - https://www.geolocation.com/
// - https://ipstack.com

// Сервис должен быть обернут в docker и опубликован на github.

// Успеешь сделать за неделю?

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/roxyash/go_ahead/pkg/config"
	_ "github.com/roxyash/go_ahead/weather_service/docs"
	"github.com/roxyash/go_ahead/weather_service/internal/response"
	"github.com/roxyash/go_ahead/weather_service/internal/service"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

type Handler struct {
	service  *service.Service
	response *response.Response
	cfg      *config.WeatherServiceConfig
}

func NewHandler(service *service.Service, response *response.Response, cfg *config.WeatherServiceConfig) *Handler {
	return &Handler{service: service, response: response, cfg: cfg}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.GET("/swagger", gin.WrapH(http.RedirectHandler("/swagger/index.html", http.StatusMovedPermanently)))
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// router.User(middleware), for auth with header. 
	router.GET("/forecast", h.forecast)

	monitoring := router.Group("/monitoring")
	{
		monitoring.GET("/healthcheck", h.healthcheck)
		monitoring.GET("/about", h.about)
	}

	return router
}
