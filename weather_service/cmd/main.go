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
	"github.com/roxyash/go_ahead/weather_service/internal/response"
	"github.com/roxyash/go_ahead/weather_service/internal/server"
	"github.com/roxyash/go_ahead/weather_service/internal/service"
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

	logger.Infof("Initialize apikeys . . .")

	if os.Getenv("geolocation_apikey") == "" {
		logger.Fatalf("error load env = geolacation_apikey, set geolacation_apikey in file weather_service/config/.env")
	}

	if os.Getenv("free_weather_apikey") == "" {
		logger.Fatalf("error load env = free_weather_apikey, set free_weather_apikey in file weather_service/config/.env")
	}

	geolocation_apikey := os.Getenv("geolocation_apikey")

	var weather_apikey string
	weatherPlan := os.Getenv("weather_plan")
	switch weatherPlan {
	case "Paid":
		if os.Getenv("paid_weather_apikey") == "" {
			logger.Fatalf("error load env = paid_weather_apikey, set paid_weather_apikey in file weather_service/config/.env")
		}
		weather_apikey = os.Getenv("paid_weather_apikey")
	case "Free":
		weather_apikey = os.Getenv("free_weather_apikey")
	default:
		weather_apikey = os.Getenv("free_weather_apikey")
	}

	logger.Infof("Initialize packages of service . . .")

	response := response.NewResponse(logger)

	services := service.NewService()

	handlers := handler.NewHandler(services, response, &config.WeatherServiceConfig{
		LocationApiKey: geolocation_apikey,
		WeatherApiKey:  weather_apikey,
		Plan: weatherPlan, 
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


// Func for initialize config with viper. 
func initConfig() error {
	viper.AddConfigPath("weather_service/config")
	viper.SetConfigName("dev")
	return viper.ReadInConfig()
}
