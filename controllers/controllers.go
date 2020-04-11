package controllers

import (
	"net/http"

	"github.com/boss959595/dbp/models"
	"github.com/labstack/echo"
)

func GetEmployees(c echo.Context) error {
	result := models.GetEmployee()
	println("foo")
	return c.JSON(http.StatusOK, result)
}
