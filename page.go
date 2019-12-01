package main

import (
	"github.com/labstack/echo"
	"go-gnuboard/controller"
	"io/ioutil"
	"net/http"
)

// IndexPage 함수는 인덱스 페이지를 보여주는 함수
// dbconfig.json 파일이 없을 경우 인스톨 페이지를 보여줌.
func IndexPage(c echo.Context) error {
	if exist := controller.FileExist(GG5.DataFolder + "dbconfig.json"); exist == true {
		// dbconfig.json이 있다면
		return c.Render(http.StatusOK, "index", "Hello Go-Gnuboard")
	} else {
		// dbconfig.json이 없다면
		return c.Render(http.StatusOK, "install", map[string]interface{}{
			"DBConfigFile": GG5.DataFolder + "dbconfig.json",
			"Version":      GG5.Version,
			"Domain":       GG5.Domain,
		})
	}
}

// InstallPage 함수는 인스톨 페이지를 보여주는 함수
func InstallPage(c echo.Context) error {
	LicenseTxt, err := ioutil.ReadFile(GG5.Path + "\\public\\LICENSE.txt")
	controller.ErrorController(err)

	return c.Render(http.StatusOK, "installer", map[string]interface{}{
		"LicenseTxt": string(LicenseTxt),
	})
}
