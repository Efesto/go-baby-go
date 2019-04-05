package openweather

import (
	"../city"

	"io/ioutil"
	"net/http"
	"os"
)

// Key returns open weather api key from environment variable OPEN_WEATHER_API_KEY
func Key() string {
	return os.Getenv("OPEN_WEATHER_API_KEY")
}

func Weather(city city.City) ([]byte, error) {
	apiKey := Key()

	url := "https://api.openweathermap.org/data/2.5/weather?q=" + city.Name + "&appid=" + apiKey
	resp, err := http.Get(url)
	if err != nil {
		defer resp.Body.Close()
		return ioutil.ReadAll(resp.Body)
	}
	return nil, err
}
