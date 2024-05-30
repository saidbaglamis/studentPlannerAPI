package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"studentPlanner/internal/database"
	models2 "studentPlanner/internal/models"
	"studentPlanner/internal/services"
)

var err error

func getStudents(c echo.Context) error {
	var students []models2.Student
	database.Db.Preload("Plans").Find(&students)
	return c.JSON(http.StatusOK, students)
}

func getStudent(c echo.Context) error {
	id := c.Param("id")
	var student models2.Student

	if result := database.Db.First(&student, id); result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "Student not found"})
	}
	database.Db.Preload("Plans").Find(&student)
	return c.JSON(http.StatusOK, student)
}

func CreateStudent(c echo.Context) error {
	student := new(models2.Student)
	if err := c.Bind(student); err != nil {
		return err
	}
	database.Db.Create(student)
	return c.JSON(http.StatusCreated, student)
}

func createPlan(c echo.Context) error {
	plan := new(models2.Plan)
	if err := c.Bind(plan); err != nil {
		return err
	}
	database.Db.Create(plan)
	return c.JSON(http.StatusCreated, plan)
}

func updatePlan(c echo.Context) error {
	id := c.Param("id")
	var plan models2.Plan
	if err := database.Db.First(&plan, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Plan not found"})
	}

	if err := c.Bind(&plan); err != nil {
		return err
	}
	database.Db.Save(&plan)
	return c.JSON(http.StatusOK, plan)
}

func deletePlan(c echo.Context) error {
	id := c.Param("id")
	if err := database.Db.Delete(&models2.Plan{}, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Plan not found"})
	}
	return c.NoContent(http.StatusNoContent)
}
func updateStudent(c echo.Context) error {
	id := c.Param("id")
	var student models2.Student
	if result := database.Db.First(&student, id); result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "Student not found"})
	}

	if err := c.Bind(&student); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid input"})
	}

	database.Db.Save(&student)
	return c.JSON(http.StatusOK, student)
}

func main() {
	e := echo.New()

	database.InitDB()

	e.POST("/student", CreateStudent)
	e.PUT("/student/:id", updateStudent)
	e.GET("/students", getStudents)
	e.GET("/student/:id", getStudent)
	e.DELETE("/student/:id", services.DeleteStudent)
	e.POST("/plans", createPlan)
	e.PUT("/plans/:id", updatePlan)
	e.DELETE("/plans/:id", deletePlan)

	e.Logger.Fatal(e.Start(":8090"))
}
