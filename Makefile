build:
	go build -o main weather_service/cmd/main.go 

go_run:build
	./main

run:
	docker-compose up 

swag:
	swag init -g ./weather_service/cmd/main.go --output weather_service/docs

.PHONY: swag build run