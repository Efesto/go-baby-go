package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.Logger.Fatal(e.Start(":1323"))

	e.GET("/greetings", greetings)
}

func greetings(context echo.Context) error {
	return context.String(http.StatusOK, "Hello, World!")
}
