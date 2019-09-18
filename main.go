package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	fmt.Println("sms-survey-demo-Back")
	//fmt.Println("TWILIO_SID:", os.Getenv("TWILIO_SID"))
	//fmt.Println("TWILIO_TOKEN:", os.Getenv("TWILIO_TOKEN"))
	if os.Getenv("TWILIO_SID") == "" || os.Getenv("TWILIO_TOKEN") == "" {
		fmt.Println("TWILIO_SID and TWILIO_TOKEN environment variables must be set.")
		fmt.Println("this can be done inside .env if using VSCode IDE")
		return
	}

	e := echo.New()

	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	e.GET("/", defaultHandler)
	e.GET("/hi", hiHandler)

	e.Logger.Fatal(e.Start(":3050"))
}

func hiHandler(c echo.Context) error {
	return c.String(http.StatusOK, "hi endpoint!\n")
}

func defaultHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Default endpoint!\n")
}
