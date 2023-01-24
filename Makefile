build:
	go build -o main weather_service/cmd/main.go 


run:build
	./main

swag:
	swag init -g ./weather_service/cmd/main.go --output weather_service/docs

.PHONY: swag build run