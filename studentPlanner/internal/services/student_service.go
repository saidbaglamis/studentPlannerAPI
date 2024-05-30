package services

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"studentPlanner/internal/database"
	"studentPlanner/internal/models"
)

func DeleteStudent(c echo.Context) error {
	id := c.Param("id")
	database.Db.Where("student_id = ?", id).Delete(&models.Plan{})
	var student models.Student
	if result := database.Db.Delete(&student, id); result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "Student not found"})
	}
	return c.NoContent(http.StatusNoContent)
}
