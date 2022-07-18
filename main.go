package main

import (
	"log"

	"github.com/adsl99801/zoo/controller"
	"github.com/adsl99801/zoo/dao"
	"github.com/adsl99801/zoo/ig"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	// postgressql "gorm.io/driver/postgres"
	// "gorm.io/gorm"
	gorm "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
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
	e.GET("/igAuth", ig.IgAuth)
	e.GET("/users", controller.GetAllUsers)
	e.POST("/users", controller.CreateUser)
	e.GET("/users/:id", controller.GetUser)
	e.PUT("/users/:id", controller.UpdateUser)
	e.DELETE("/users/:id", controller.DeleteUser)
	e.Static("/", "public")
	e.File("/", "public/index.html")
	e.Logger.Fatal(e.Start(":8080"))
}
