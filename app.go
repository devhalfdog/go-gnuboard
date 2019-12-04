/*
	PHP 그누보드를 Golang 으로 이식하기 위한 프로젝트
*/

package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"go-gnuboard/conf"
)

var GG5 *conf.GG5S

func main() {
	GG5 = conf.NewConfig()

	e := echo.New()
	if GG5.Debug {
		e.Use(middleware.Logger())
	}

	router(e)

	e.Logger.Fatal(e.Start(GG5.Port))
}
