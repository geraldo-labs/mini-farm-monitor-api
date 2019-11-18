package main

import (
	"net/http"

	v "github.com/geraldo-labs/farm-monitor/version"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.JSONPretty(http.StatusOK, v.CurrentVersion, "  ")
	})
	e.Logger.Fatal(e.Start(":1323"))
}