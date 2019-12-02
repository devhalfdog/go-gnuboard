package main

import (
	"github.com/labstack/echo"
	"net/http"
	"time"
)

// TODO -- Header 작성

// InstallProcessHeader 미들웨어 함수는 Header를 추가합니다.
func InstallProcessHeader(next echo.HandlerFunc) echo.HandlerFunc {
	// TODO --- Header가 일정 페이지에서만 작동하게 하는 거 찾을 것.
	return func(c echo.Context) error {
		c.Response().Header().Set("Expires", "0")
		c.Response().Header().Set(echo.HeaderLastModified, time.Now().Format("Mon, 02 Dec 2019 09:56:15 GMT")) // Mon, 02 Dec 2019 09:56:15 GMT
		c.Response().Header().Set("Cache-Control", "no-store, no-cache, must-revalidate")
		c.Response().Header().Set("Cache-Control", "pre-check=0, post-check=0, max-age=0")
		c.Response().Header().Set("Pragma", "no-cache")
		c.Response().Header().Set(echo.HeaderContentType, "text/html; charset=utf-8")
		c.Response().Header().Set("X-Robots-Tag", "noindex")

		return next(c)
	}
}

// TODO - 인스톨 프로세스 작성할 것.
func InstallProcess(c echo.Context) error {
	c.Response().Header().Set("Expires", "0")
	c.Response().Header().Set(echo.HeaderLastModified, time.Now().Format("Mon, 02 Dec 2019 09:56:15 GMT")) // Mon, 02 Dec 2019 09:56:15 GMT
	c.Response().Header().Set("Cache-Control", "no-store, no-cache, must-revalidate")
	c.Response().Header().Set("Cache-Control", "pre-check=0, post-check=0, max-age=0")
	c.Response().Header().Set("Pragma", "no-cache")
	c.Response().Header().Set(echo.HeaderContentType, "text/html; charset=utf-8")
	c.Response().Header().Set("X-Robots-Tag", "noindex")
	c.Response().WriteHeader(http.StatusOK)

	if agree := c.FormValue("agree"); agree == "" {
		return c.HTML(http.StatusOK, `<!doctype html><html><head><meta charset="UTF-8"><title>`)
	}

	return c.String(http.StatusOK, "Test")
}
