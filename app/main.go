package main

import (
	"github.com/dwin/hashify/app/controller"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {

	e := Router()
	e.Logger.Fatal(e.Start(":1313"))
}

func Router() *echo.Echo {
	e := echo.New()
	// Limit Request Body to 25 MB
	e.Use(middleware.BodyLimit("25M"))

	// Routes
	e.GET("/status", controller.GetStatus)
	// SHA Routes
	e.POST("/sha1", controller.HashSHA1)
	e.GET("/sha1", controller.HashSHA1)
	// HighwayHash Routes
	e.POST("/highway", controller.HashHighwayHash) // default 256 function
	e.GET("/highway", controller.HashHighwayHash)
	e.POST("/highway64", controller.HashHighwayHash64)
	e.GET("/highway64", controller.HashHighwayHash64)
	e.POST("/highway128", controller.HashHighwayHash128)
	e.GET("/highway128", controller.HashHighwayHash128)

	return e
}
