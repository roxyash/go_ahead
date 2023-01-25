# weather_service

## This service contains 3 methods.
### `[GET] /forecast`
The method accepts query parameters:
- date — forecast day date (current day by default);
- city — city (by default, the service should calculate the client's current location).

In response, the service gives the average temperature per day in degrees and the location.

### `[GET] /monitoring/about`
Service Information
### `[GET] /monitoring/healthcheck`
Method for checking the life of a service


## Info

Before starting the application, you need to get `YOUR` keys in the services:
1. https://openweathermap.org/api (To determine the weather)
2. https://ipstack.com/ (For geolocation)
3. https://www.weatherapi.com/ (To determine weather) `[FREE]`

After that, add the keys to the .env file along the path `weather_service/config /.env`

The `.env` file was not created because I didn't push it to the repository.

You need to create it and add two fields there:

```powershell
paid_weather_apikey=your_api_key
free_weather_apikey=your_api_key
geolocation_apikey=your_apikey
weather_plan="Free" or "Paid" default was as Free
```

## Installation

#### Use Makefile for running service
```powershell
make run 
```

After running the `make run` command in the terminal, the `docker-compose up` command will be executed, which in turn will start the weather_service in your Docker.

Before starting, make sure Docker is running :smiley:

After assembly, you can start using the application.

Open the page - `http://localhost:8000/swagger`, to access the API documentation.


