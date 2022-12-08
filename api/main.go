package main

import (
	"api/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()

	// cors policy
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	// / route for test
	e.POST("/", func(c echo.Context) error {
		return handlers.SimpleFunc(c)
	})

	port := "8080"
	e.Logger.Fatal(e.Start(":" + port))
}
