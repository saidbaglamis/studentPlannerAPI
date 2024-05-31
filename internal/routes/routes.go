package routes

import (
	"github.com/labstack/echo/v4"
	"studentPlanner/internal/handlers"
)

func InitAllRoutes() {
	e := echo.New()
	//student handlerlar
	e.GET("/student", handlers.GetAllStudents)
	e.GET("/student/:id", handlers.GetStudent)
	e.POST("/student", handlers.CreateStudent)
	e.PUT("/student/:id", handlers.UpdateStudent)
	e.DELETE("/student/:id", handlers.DeleteStudent)

	//plan handlerlar
	e.GET("/plan/:id", handlers.GetPlan)
	e.GET("/weekly/:id", handlers.GetWeeklyPlans)
	e.GET("/monthly/:id", handlers.GetMonthlyPlans)
	e.POST("/plan", handlers.CreatePlan)
	e.PUT("/plan/:id", handlers.UpdatePlan)
	e.DELETE("/plan/:id", handlers.DeletePlan)

	e.Logger.Fatal(e.Start(":8050"))
}
