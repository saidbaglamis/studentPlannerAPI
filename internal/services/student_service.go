package services

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"studentPlanner/internal/database"
	"studentPlanner/internal/models"
)

func GetStudents(c echo.Context) error {
	var students []models.Student
	database.Db.Preload("Plans").Find(&students)
	return c.JSON(http.StatusOK, students)
}

func GetStudent(c echo.Context) error {
	id := c.Param("id")
	var student models.Student

	if result := database.Db.First(&student, id); result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "Student not found"})
	}
	database.Db.Preload("Plans").Find(&student)
	return c.JSON(http.StatusOK, student)
}

func CreateStudent(c echo.Context) error {
	student := new(models.Student)
	if err := c.Bind(student); err != nil {
		return err
	}
	database.Db.Create(student)
	return c.JSON(http.StatusCreated, student)
}

func UpdateStudent(c echo.Context) error {
	id := c.Param("id")
	var student models.Student
	if result := database.Db.First(&student, id); result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "Student not found"})
	}

	if err := c.Bind(&student); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid input"})
	}

	database.Db.Save(&student)
	return c.JSON(http.StatusOK, student)
}

func DeleteStudent(c echo.Context) error {
	id := c.Param("id")
	database.Db.Where("student_id = ?", id).Delete(&models.Plan{})
	var student models.Student
	if result := database.Db.Delete(&student, id); result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "Student not found"})
	}
	return c.NoContent(http.StatusNoContent)
}
