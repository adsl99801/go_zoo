package main

import (
	"github.com/adsl99801/zoo/ig"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/igAuth", ig.IgAuth)
	e.Static("/", "public")
	e.File("/", "public/index.html")
	e.Logger.Fatal(e.Start(":8080"))
}
