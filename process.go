package main

import (
	"github.com/labstack/echo"
	"net/http"
)

func InstallProcess(c echo.Context) error {
	return c.String(http.StatusOK, "Test")
}
