package main

import (
	"api-playground/database"
	"api-playground/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	database.ConnectDatabase()

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.POST("/CreateNewUser", handlers.CreateNewUser)
	e.GET("/GetAllUsers", handlers.GetAllUsers)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
