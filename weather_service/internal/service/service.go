package service

import "time"

//go:generate mockgen -source=service.go -destination=mocks/mock.go
type Weather interface {
	GetOpenWeatherFree(apiKey, city string, date time.Time) (string, error)
	GetWeatherFree(apiKey, city string, date time.Time) (string, error)
	GetOpenWeatherPaid(apiKey, city string, date time.Time) (string, error)
}

type Location interface {
	Get(ip, apiKey string) (string, error)
}

type Ip interface {
	Get() (string, error)
}

type Service struct {
	Ip
	Weather
	Location
}

func NewService() *Service {
	return &Service{
		Ip:       NewIpService(),
		Weather:  NewWeatherService(),
		Location: NewLocationService(),
	}

}
