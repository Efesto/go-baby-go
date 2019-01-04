package main

import (
	"net/http"

	"../pkg/city"
	"../pkg/openweather"
	"github.com/labstack/echo"
	"github.com/subosito/gotenv"
)

func main() {
	gotenv.Load()

	e := echo.New()

	e.GET("/greetings", greetings)
	e.GET("/weather/:city", weather)

	e.Logger.Fatal(e.Start(":5000"))
}

func greetings(context echo.Context) error {
	return context.String(http.StatusOK, "Hello, World!")
}

func weather(context echo.Context) error {
	city := city.City{Name: context.Param("city")}

	body, _ := openweather.Weather(city)
	return context.String(http.StatusOK, string(body))
}
