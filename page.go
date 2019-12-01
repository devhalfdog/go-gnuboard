package main

import (
	"github.com/labstack/echo"
	"go-gnuboard/controller"
	"net/http"
)

func IndexPage(c echo.Context) error {
	if _, err := controller.FileExist(GG5.DataFolder + "dbconfig.json"); err == nil {
		// dbconfig.json이 없다면
		return c.Render(http.StatusOK, "install", map[string]interface{}{
			"DBConfigFile": GG5.DataFolder + "dbconfig.json",
			"Version":      GG5.Version,
			"Domain":       GG5.Domain,
		})
	} else {
		// dbconfig.json이 있다면
		return c.Render(http.StatusOK, "index", "Hello Go-Gnuboard")
	}
}

func InstallPage(c echo.Context) error {
	return c.Render(http.StatusOK, "installer", "")
}
