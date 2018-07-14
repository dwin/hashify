package controller

import "github.com/labstack/echo"

func GetIndex(c echo.Context) error {
	return c.Redirect(302, "https://hashify.net")
}
