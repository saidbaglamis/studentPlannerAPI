package services

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"studentPlanner/internal/database"
	"studentPlanner/internal/models"
)

func CreatePlan(c echo.Context) error {
	plan := new(models.Plan)
	if err := c.Bind(plan); err != nil {
		return err
	}
	database.Db.Create(plan)
	return c.JSON(http.StatusCreated, plan)
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
