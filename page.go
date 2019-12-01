package main

import (
	"github.com/labstack/echo"
	"net/http"
)

func IndexPage(c echo.Context) error {
	return c.Render(http.StatusOK, "index", "Hello Go-Gnuboard")
}
