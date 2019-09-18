package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	fmt.Println("sms-survey-demo-Back")
	e := echo.New()

	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		//AllowOrigins: []string{"*", "http://home.mevise.com:3000", "localhost"},
		AllowOrigins: []string{"*"},
		//AllowMethods: []string{echo.GET, echo.HEAD, echo.POST},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	e.GET("/", defaultHandler)
	e.GET("/hi", hiHandler)

	e.Logger.Fatal(e.Start(":3051"))
}

func hiHandler(c echo.Context) error {
	//fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))

	return c.String(http.StatusOK, "hi endpoint!\n")
}

func defaultHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Default endpoint!\n")
}
