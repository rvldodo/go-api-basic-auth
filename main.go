package main

import (
	"fmt"
	"go-api-basic-auth/src/config"
	"go-api-basic-auth/src/model"
	"go-api-basic-auth/src/router"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	db, err := config.ConnectDB()
	if err != nil {
		panic(err)
	}

	fmt.Println("DB connected")
	db.AutoMigrate(&model.Student{})

	router.StudentRouter(e)

	e.Logger.Fatal(e.Start(":8080"))
}