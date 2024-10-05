package controllers

import (
	"net/http"

	"github.com/reusmana/car-rental-go/config"
	"github.com/reusmana/car-rental-go/models"
	"github.com/reusmana/car-rental-go/utils"

	"github.com/labstack/echo/v4"
)

func GetCars(c echo.Context) error {
	var bookings []models.Car
	config.DB.Find(&bookings)
	return utils.JSONResponse(c, http.StatusOK, "success", bookings)
}

func GetCarById(c echo.Context) error {
	id := c.Param("id")
	var car models.Car

	if result := config.DB.First(&car, id); result.Error != nil {
		return utils.JSONResponse(c, http.StatusNotFound, "car not found", nil)
	}
	return utils.JSONResponse(c, http.StatusOK, "success", car)

}

func CreateCar(c echo.Context) error {
	car := new(models.Car)
	tx := config.DB.Begin()

	if err := c.Bind(car); err != nil {
		return utils.JSONResponse(c, http.StatusBadRequest, err.Error(), nil)
	}

	if result := tx.Create(&car); result.Error != nil {
		tx.Rollback()
		return utils.JSONResponse(c, http.StatusInternalServerError, result.Error.Error(), nil)
	}
	tx.Commit()
	return utils.JSONResponse(c, http.StatusCreated, "success created cars", car)
}

func UpdateCar(c echo.Context) error {
	id := c.Param("id")
	var car models.Car
	tx := config.DB.Begin()

	if err := tx.First(&car, id).Error; err != nil {
		return utils.JSONResponse(c, http.StatusNotFound, "Car not found", nil)
	}

	if err := c.Bind(&car); err != nil {
		return utils.JSONResponse(c, http.StatusBadRequest, "Invalid input", nil)
	}

	if err := tx.Save(&car).Error; err != nil {
		tx.Rollback()
		return utils.JSONResponse(c, http.StatusInternalServerError, "Failed to update car", nil)
	}

	tx.Commit()

	return utils.JSONResponse(c, http.StatusOK, "Car updated successfully", car)

}

func DeleteCar(c echo.Context) error {
	id := c.Param("id")
	var car models.Car
	tx := config.DB.Begin()

	if err := tx.First(&car, id).Error; err != nil {
		return utils.JSONResponse(c, http.StatusNotFound, "Car not found", nil)
	}

	if err := tx.Delete(&car).Error; err != nil {
		tx.Rollback()
		return utils.JSONResponse(c, http.StatusInternalServerError, "Failed to delete car", nil)
	}

	tx.Commit()

	return utils.JSONResponse(c, http.StatusOK, "Car deleted successfully", nil)

}
