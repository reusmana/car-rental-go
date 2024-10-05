package controllers

import (
	"net/http"

	"github.com/reusmana/car-rental-go/config"
	"github.com/reusmana/car-rental-go/models"
	"github.com/reusmana/car-rental-go/utils"

	"github.com/labstack/echo/v4"
)

func GetDriver(c echo.Context) error {
	var driver []models.Driver
	config.DB.Find(&driver)

	return utils.JSONResponse(c, http.StatusOK, "success", driver)
}

func CreateDriver(c echo.Context) error {
	driver := new(models.Driver)
	tx := config.DB.Begin()

	if err := c.Bind(driver); err != nil {
		return utils.JSONResponse(c, http.StatusBadRequest, err.Error(), nil)
	}

	if result := tx.Create(&driver); result.Error != nil {
		tx.Rollback()
		return utils.JSONResponse(c, http.StatusInternalServerError, result.Error.Error(), nil)
	}
	tx.Commit()
	return utils.JSONResponse(c, http.StatusCreated, "success created drivers", driver)
}
