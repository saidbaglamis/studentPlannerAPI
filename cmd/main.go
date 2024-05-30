package main

import (
	"github.com/labstack/echo/v4"
	"studentPlanner/internal/database"
	"studentPlanner/internal/services"
)

func main() {
	e := echo.New()

	database.InitDB()
	//student handlerlari
	e.POST("/student", services.CreateStudent)
	e.PUT("/student/:id", services.UpdateStudent)
	e.GET("/students", services.GetStudents)
	e.GET("/student/:id", services.GetStudent)
	e.DELETE("/student/:id", services.DeleteStudent)

	//plan handlerları
	e.POST("/plans", services.CreatePlan)
	e.PUT("/plans/:id", services.UpdatePlan)
	e.DELETE("/plans/:id", services.DeletePlan)

	e.Logger.Fatal(e.Start(":8050"))
}
