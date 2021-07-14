package main

import (
	"github.com/adsl99801/zoo/controller"
	"github.com/adsl99801/zoo/ig"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/static", "assets")
	e.GET("/igAuth", ig.IgAuth)
	e.GET("/", account.Home)
	// servers other static files
	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
