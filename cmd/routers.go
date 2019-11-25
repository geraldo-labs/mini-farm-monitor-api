package main

import (
	"net/http"
	"github.com/labstack/echo"
	v "github.com/geraldo-labs/farm-monitor/version"
)

func Home(c echo.Context) error {
	return c.JSONPretty(http.StatusOK, v.CurrentVersion, "  ")
}