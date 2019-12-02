package main

import (
	"github.com/labstack/echo"
	"net/http"
)

// TODO - 인스톨 프로세스 작성할 것.
func InstallProcess(c echo.Context) error {
	return c.String(http.StatusOK, "Test")
}
