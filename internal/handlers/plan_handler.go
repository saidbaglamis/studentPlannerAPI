package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"studentPlanner/internal/database"
	"studentPlanner/internal/models"
	"time"
)

func GetWeeklyPlans(c echo.Context) error {
	var plans []models.Plan
	studentID := c.Param("id")
	startOfWeek := time.Now().UTC().Truncate(24*time.Hour).AddDate(0, 0, -int(time.Now().Weekday()))
	endOfWeek := startOfWeek.AddDate(0, 0, 7)
	database.Db.Where("student_id = ? AND start_date >= ? AND end_date < ?", studentID, startOfWeek, endOfWeek).Find(&plans)
	return c.JSON(http.StatusOK, plans)
}

func GetMonthlyPlans(c echo.Context) error {
	var plans []models.Plan
	studentID := c.Param("id")
	now := time.Now().UTC()
	startOfMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
	endOfMonth := startOfMonth.AddDate(0, 1, 0)
	database.Db.Where("student_id = ? AND start_date >= ? AND end_date < ?", studentID, startOfMonth, endOfMonth).Find(&plans)
	return c.JSON(http.StatusOK, plans)
}

func CreatePlan(c echo.Context) error {
	plan := new(models.Plan)
	if err := c.Bind(plan); err != nil {
		return err
	}
	var overlappingPlans []models.Plan
	database.Db.Where("student_id = ? AND start_date < ? AND end_date > ?", plan.StudentID, plan.EndDate, plan.StartDate).Find(&overlappingPlans)
	if len(overlappingPlans) > 0 {
		return c.JSON(http.StatusConflict, map[string]string{"message": "Plan conflicts with an existing plan."})
	}
	database.Db.Create(plan)
	return c.JSON(http.StatusCreated, plan)
}
func GetPlan(c echo.Context) error {
	plan := new(models.Plan)
	id := c.Param("id")
	if err := database.Db.First(&plan, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Plan not found"})
	}
	return c.JSON(http.StatusOK, plan)
}

func UpdatePlan(c echo.Context) error {
	id := c.Param("id")
	var plan models.Plan
	if err := database.Db.First(&plan, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Plan not found"})
	}

	if err := c.Bind(&plan); err != nil {
		return err
	}
	database.Db.Save(&plan)
	return c.JSON(http.StatusOK, plan)
}

func DeletePlan(c echo.Context) error {
	id := c.Param("id")
	if err := database.Db.Delete(&models.Plan{}, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Plan not found"})
	}
	return c.NoContent(http.StatusNoContent)
}
