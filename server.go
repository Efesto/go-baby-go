package main

import (
	"io/ioutil"
	"net/http"
	"os"

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
	city := context.Param("city")
	apiKey := openWeatherAPIKey()

	url := "https://api.openweathermap.org/data/2.5/weather?q=" + city + "&appid=" + apiKey
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	return context.String(http.StatusOK, string(body))
}

func openWeatherAPIKey() string {
	return os.Getenv("OPEN_WEATHER_API_KEY")
}
