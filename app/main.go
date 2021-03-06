package main

import (
	"net/http"
	"time"

	"github.com/dwin/hashify/app/controller"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := Router()
	controller.StartTime = time.Now()
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 1 << 10, // 1 KB
	}))
	e.Logger.Fatal(e.Start(":1313"))
}

func Router() *echo.Echo {
	e := echo.New()

	e.Pre(middleware.BodyLimit("10M"))
	// Limit Querystring length for "value"
	e.Pre(QueryLength)

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"localhost", "https://hashify.net", "*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, "X-Hashify-Key", "X-Hashify-Process"},
	}))

	// Routes
	e.GET("/", controller.GetIndex)
	e.GET("/status", controller.GetStatus)
	e.GET("/methods", controller.ListMethods)
	e.GET("/keygen/:length", controller.KeyGen)
	// Limit Request Body to 10 MB
	limit := e.Group("", middleware.BodyLimit("10M"))
	h := limit.Group("/hash")
	h.GET("/:algo/:format", controller.ComputeHash)
	h.POST("/:algo/:format", controller.ComputeHash)
	// MD5 Route
	/*
		e.POST("/md5", controller.HashMD5)
		e.GET("/md5", controller.HashMD5)
	*/
	// HighwayHash Routes
	h.POST("/highway", controller.HashHighwayHash) // default 256 function
	h.GET("/highway", controller.HashHighwayHash)
	h.POST("/highway64", controller.HashHighwayHash64)
	h.GET("/highway64", controller.HashHighwayHash64)
	h.POST("/highway128", controller.HashHighwayHash128)
	h.GET("/highway128", controller.HashHighwayHash128)
	// SHA Routes
	/*
		e.POST("/sha1", controller.HashSHA1)
		e.GET("/sha1", controller.HashSHA1)
		e.POST("/sha256", controller.HashSHA256)
		e.GET("/sha256", controller.HashSHA256)
		e.POST("/sha384", controller.HashSHA384)
		e.GET("/sha384", controller.HashSHA384)
		e.POST("/sha512", controller.HashSHA512)
		e.GET("/sha512", controller.HashSHA512)
		// SHA3 Routes
		e.POST("/sha3-256", controller.HashSHA3_256)
		e.GET("/sha3-256", controller.HashSHA3_256)
		e.POST("/sha3-384", controller.HashSHA3_384)
		e.GET("/sha3-384", controller.HashSHA3_384)
		e.POST("/sha3-512", controller.HashSHA3_512)
		e.GET("/sha3-512", controller.HashSHA3_512)
	*/
	return e
}

func QueryLength(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if len(c.QueryParam("value")) > 1500 {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "value parameter too long, must not be more than 1500 characters",
			})
		}
		return next(c)
	}
}
