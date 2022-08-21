package main

import (
	"log"

	"github.com/adsl99801/zoo/controller"
	"github.com/adsl99801/zoo/dao"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	// postgressql "gorm.io/driver/postgres"
	// "gorm.io/gorm"
	gorm "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)
const (
	apiUrlStr = "api"
)
func main() {
	
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	
	dsn := "postgres://pfuser:pfuser@192.168.0.104/pf" 
	// db, err := gorm.Open(postgressql.Open(dsn), &gorm.Config{})
	// db, err := gorm.Open("postgres", "pfuser:pfuser@192.168.0.104/pf?charset=utf8&parseTime=True&loc=Local")
	db, err := gorm.Open("postgres", dsn)
	if err != nil {
        log.Fatalln(err)
    }
	dao.DB = db;
	apiPost(e,"/Login",controller.Login);
	e.Static("/", "public")
	e.File("/", "public/index.html")
	e.Logger.Fatal(e.Start(":8080"))
}
func apiPost(e *echo.Echo,urlStr  string, h echo.HandlerFunc){
	e.POST(apiUrlStr+urlStr, h)
}
