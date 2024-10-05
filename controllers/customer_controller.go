package controllers

import (
	"net/http"

	"github.com/reusmana/car-rental-go/config"
	"github.com/reusmana/car-rental-go/models"
	"github.com/reusmana/car-rental-go/utils"

	"github.com/labstack/echo/v4"
)

func GetCustomer(c echo.Context) error {
	var customer []models.Customer
	config.DB.Find(&customer)

	return utils.JSONResponse(c, http.StatusOK, "success", customer)
}

func GetCustomerById(c echo.Context) error {
	id := c.Param("id")
	var customer models.Customer

	if result := config.DB.First(&customer, id); result.Error != nil {
		return utils.JSONResponse(c, http.StatusNotFound, "customer not found", nil)
	}
	return utils.JSONResponse(c, http.StatusOK, "success", customer)

}

func CreateCustomer(c echo.Context) error {
	customer := new(models.Customer)
	tx := config.DB.Begin()

	if err := c.Bind(customer); err != nil {
		return utils.JSONResponse(c, http.StatusBadRequest, err.Error(), nil)
	}

	if result := tx.Create(&customer); result.Error != nil {
		tx.Rollback()
		return utils.JSONResponse(c, http.StatusInternalServerError, result.Error.Error(), nil)
	}
	tx.Commit()
	return utils.JSONResponse(c, http.StatusCreated, "success created customers", customer)
}

func UpdateCustomer(c echo.Context) error {
	id := c.Param("id")
	var customer models.Customer
	tx := config.DB.Begin()

	if err := tx.First(&customer, id).Error; err != nil {
		return utils.JSONResponse(c, http.StatusNotFound, "customer not found", nil)
	}

	if err := c.Bind(&customer); err != nil {
		return utils.JSONResponse(c, http.StatusBadRequest, "Invalid input", nil)
	}

	if err := tx.Save(&customer).Error; err != nil {
		tx.Rollback()
		return utils.JSONResponse(c, http.StatusInternalServerError, "Failed to update customer", nil)
	}

	tx.Commit()

	return utils.JSONResponse(c, http.StatusOK, "customer updated successfully", customer)

}

func DeleteCustomer(c echo.Context) error {
	id := c.Param("id")
	var customer models.Customer
	tx := config.DB.Begin()

	if err := tx.First(&customer, id).Error; err != nil {
		return utils.JSONResponse(c, http.StatusNotFound, "customer not found", nil)
	}

	if err := tx.Delete(&customer).Error; err != nil {
		tx.Rollback()
		return utils.JSONResponse(c, http.StatusInternalServerError, "Failed to delete customer", nil)
	}

	tx.Commit()

	return utils.JSONResponse(c, http.StatusOK, "customer deleted successfully", nil)

}
