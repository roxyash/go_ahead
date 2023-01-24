package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/roxyash/go_ahead/pkg/config"
	"github.com/roxyash/go_ahead/pkg/zaplogger"
	"github.com/roxyash/go_ahead/weather_service/internal/handler"
	"github.com/roxyash/go_ahead/weather_service/internal/server"
	"github.com/roxyash/go_ahead/weather_service/internal/service"
	"github.com/roxyash/go_ahead/weather_service/internal/response"
	"github.com/spf13/viper"
)


// @title Weather service 
// @version 1.0
// @description Simple weather service
// @termsOfService *
// @contact.name *
// @contact.url *
// @contact.email *
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8000
// @BasePath /
func main() {
	fmt.Println("Initialize config . . .")
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	fmt.Println("Initialize zaplogger . . .")
	logger := zaplogger.NewZapLogger(viper.GetString("app.logPath"), "")

	logger.Infof("Initialize env files. . .")
	if err := godotenv.Load(viper.GetString("app.envPath")); err != nil {
		logger.Fatalf("error loading env variables: %s", err.Error())
	}

	logger.Infof("Initialize packages of service . . .")

	response := response.NewResponse(logger)

	services := service.NewService()

	handlers := handler.NewHandler(services, response, &config.WeatherServiceConfig{
		LocationApiKey: os.Getenv("geolocation_apikey"),
		WeatherApiKey:  os.Getenv("openweathermap_apikey"),
	})

	srv := new(server.Server)

	go func() {
		if err := srv.Run(viper.GetString("app.port"), handlers.InitRoutes()); err != nil {
			logger.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	logger.Infof("weather_service on http server Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logger.Infof("weather_service on http server ShuttingDown")
	if err := srv.Shutdown(context.Background()); err != nil {
		logger.Fatalf("error occured on server shutting down :%s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("weather_service/config")
	viper.SetConfigName("dev")
	return viper.ReadInConfig()
}
