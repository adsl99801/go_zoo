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
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
        Root:   "web/dist/puppysfun/",
        Index: "index.html",
        Browse: true,
        HTML5:  true,
    }))
	e.Static("/", "web/dist/puppysfun/")
	e.File("/*", "web/dist/puppysfun/index.html")
	e.Logger.Fatal(e.Start(":8080"))
}
