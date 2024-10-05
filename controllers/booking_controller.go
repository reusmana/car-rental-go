package controllers

import (
	"errors"
	"net/http"

	"github.com/reusmana/car-rental-go/config"
	"github.com/reusmana/car-rental-go/models"
	"github.com/reusmana/car-rental-go/utils"

	"time"

	"github.com/labstack/echo/v4"
)

type BookingCarDriverId struct {
	CarID    uint  `json:"car_id"`
	DriverID *uint `json:"driver_id"`
}

func GetBookings(c echo.Context) error {
	var bookings []models.Booking
	config.DB.Find(&bookings)

	return utils.JSONResponse(c, http.StatusOK, "success", bookings)
}

func GetBookingById(c echo.Context) error {
	id := c.Param("id")
	var Booking models.Booking

	if result := config.DB.First(&Booking, id); result.Error != nil {
		return utils.JSONResponse(c, http.StatusNotFound, "Booking not found", nil)
	}
	return utils.JSONResponse(c, http.StatusOK, "success", Booking)

}

func CreateBooking(c echo.Context) error {
	booking := new(models.Booking)
	tx := config.DB.Begin()

	if err := c.Bind(booking); err != nil {
		return utils.JSONResponse(c, http.StatusBadRequest, err.Error(), nil)
	}

	var car models.Car
	if result := config.DB.Where("availability = ?", true).First(&car, booking.CarID); result.Error != nil {
		return utils.JSONResponse(c, http.StatusNotFound, "Car not found", nil)
	}
	if !car.Availability {
		return utils.JSONResponse(c, http.StatusNotFound, "Car not available", nil)
	}

	var customer models.Customer
	if result := config.DB.First(&customer, booking.CustomerID); result.Error != nil {
		return utils.JSONResponse(c, http.StatusNotFound, "Customer not found", nil)
	}

	daysOfRent, err := UtilsGetDaysOfRent(booking.StartDate, booking.EndDate)
	if err != nil {
		return utils.JSONResponse(c, http.StatusBadRequest, err.Error(), nil)
	}

	totalCost := float64(daysOfRent) * car.DailyRent

	booking.TotalCost = totalCost
	booking.DayOfRent = daysOfRent
	totalDiscount := calculateDiscount(daysOfRent, car.DailyRent, customer.MembershipID)
	booking.MembershipDiscount = totalDiscount

	if booking.RentType == "with_driver" {
		var driver models.Driver
		if err := tx.First(&driver, booking.DriverID).Error; err != nil {
			return utils.JSONResponse(c, http.StatusNotFound, "Driver not found", nil)
		}

		booking.DriverIncentive = calculateDriverIncentive(daysOfRent, car.DailyRent)
		driver.Incentive += booking.DriverIncentive

		if result := tx.Save(&driver); result.Error != nil {
			tx.Rollback()
			return utils.JSONResponse(c, http.StatusInternalServerError, result.Error.Error(), nil)
		}
	}

	// Calculate total cost
	booking.TotalCost = (float64(daysOfRent) * car.DailyRent) - totalDiscount

	booking.Status = true
	car.Availability = false

	if result := tx.Create(&booking); result.Error != nil {
		tx.Rollback()
		return utils.JSONResponse(c, http.StatusInternalServerError, result.Error.Error(), nil)
	}
	if err := tx.Save(&car).Error; err != nil {
		tx.Rollback()
		return utils.JSONResponse(c, http.StatusInternalServerError, "Failed to update car available", nil)
	}
	tx.Commit()
	return utils.JSONResponse(c, http.StatusCreated, "success created bookings", booking)
}

func UpdateBooking(c echo.Context) error {
	id := c.Param("id")
	var booking models.Booking
	tx := config.DB.Begin()

	if err := tx.First(&booking, id).Error; err != nil {
		return utils.JSONResponse(c, http.StatusNotFound, "booking not found", nil)
	}

	response := BookingCarDriverId{
		CarID:    booking.CarID,
		DriverID: booking.DriverID,
	}

	if err := c.Bind(&booking); err != nil {
		return utils.JSONResponse(c, http.StatusBadRequest, "Invalid input", nil)
	}

	var car models.Car

	if result := config.DB.First(&car, booking.CarID); result.Error != nil {
		return utils.JSONResponse(c, http.StatusNotFound, "Car not found", nil)
	}

	if response.CarID != booking.CarID {
		if !car.Availability {
			return utils.JSONResponse(c, http.StatusNotFound, "Car not available", nil)
		}
	}

	var customer models.Customer
	if result := tx.First(&customer, booking.CustomerID); result.Error != nil {
		return utils.JSONResponse(c, http.StatusNotFound, "Customer not found", nil)
	}

	daysOfRent, err := UtilsGetDaysOfRent(booking.StartDate, booking.EndDate)
	if err != nil {
		return utils.JSONResponse(c, http.StatusBadRequest, err.Error(), nil)
	}

	totalCost := float64(daysOfRent) * car.DailyRent

	booking.TotalCost = totalCost
	booking.DayOfRent = daysOfRent

	totalDiscount := calculateDiscount(daysOfRent, car.DailyRent, customer.MembershipID)
	booking.MembershipDiscount = totalDiscount

	if booking.Status {

		if booking.RentType == "with_driver" {
			var driver models.Driver
			if response.DriverID != nil && (response.DriverID != booking.DriverID) {

				if err := tx.First(&driver, response.DriverID).Error; err != nil {
					return utils.JSONResponse(c, http.StatusNotFound, "Driver not found", nil)
				}
				driver.Incentive -= booking.DriverIncentive
				booking.DriverIncentive = 0
				if result := tx.Save(&driver); result.Error != nil {
					tx.Rollback()
					return utils.JSONResponse(c, http.StatusInternalServerError, result.Error.Error(), nil)
				}

			} else {
				if response.DriverID == nil {
					if err := tx.First(&driver, booking.DriverID).Error; err != nil {
						return utils.JSONResponse(c, http.StatusNotFound, "Driver not found", nil)
					}

					booking.DriverIncentive = calculateDriverIncentive(daysOfRent, car.DailyRent)
					driver.Incentive += booking.DriverIncentive
					if result := tx.Save(&driver); result.Error != nil {
						tx.Rollback()
						return utils.JSONResponse(c, http.StatusInternalServerError, result.Error.Error(), nil)
					}
				}

			}

		} else {

			if response.DriverID != nil {
				var driver models.Driver
				if err := tx.First(&driver, response.DriverID).Error; err != nil {
					return utils.JSONResponse(c, http.StatusNotFound, "Driver not found", nil)
				}
				driver.Incentive -= booking.DriverIncentive
				if result := tx.Save(&driver); result.Error != nil {
					tx.Rollback()
					return utils.JSONResponse(c, http.StatusInternalServerError, result.Error.Error(), nil)
				}
				booking.DriverIncentive = 0
				booking.DriverID = nil
			}
		}
	}

	// Calculate total cost
	booking.TotalCost = (float64(daysOfRent) * car.DailyRent) - totalDiscount

	if booking.Status {
		car.Availability = false
		if err := tx.Save(&car).Error; err != nil {
			tx.Rollback()
			return utils.JSONResponse(c, http.StatusInternalServerError, "Failed to update booking", nil)
		}
	} else {
		car.Availability = true
		if err := tx.Save(&car).Error; err != nil {
			tx.Rollback()
			return utils.JSONResponse(c, http.StatusInternalServerError, "Failed to update booking", nil)
		}

	}

	if response.CarID != booking.CarID {
		car.Availability = false
		if err := tx.Save(&car).Error; err != nil {
			tx.Rollback()
			return utils.JSONResponse(c, http.StatusInternalServerError, "Failed to update booking", nil)
		}

		var oldCar models.Car
		if err := config.DB.First(&oldCar, response.CarID).Error; err != nil {
			return utils.JSONResponse(c, http.StatusInternalServerError, "Car not found", nil)
		}

		oldCar.Availability = true
		if err := tx.Save(&oldCar).Error; err != nil {
			tx.Rollback()
			return utils.JSONResponse(c, http.StatusInternalServerError, "Failed to update booking", nil)
		}
	}

	if err := tx.Save(&booking).Error; err != nil {
		tx.Rollback()
		return utils.JSONResponse(c, http.StatusInternalServerError, "Failed to update booking", nil)
	}

	tx.Commit()

	return utils.JSONResponse(c, http.StatusOK, "booking updated successfully", booking)

}

func DeleteBooking(c echo.Context) error {
	id := c.Param("id")
	var booking models.Booking
	tx := config.DB.Begin()

	if err := tx.First(&booking, id).Error; err != nil {
		return utils.JSONResponse(c, http.StatusNotFound, "booking not found", nil)
	}

	if err := tx.Delete(&booking).Error; err != nil {
		tx.Rollback()
		return utils.JSONResponse(c, http.StatusInternalServerError, "Failed to delete booking", nil)
	}

	tx.Commit()
	return utils.JSONResponse(c, http.StatusOK, "booking deleted successfully", nil)

}

func UtilsGetDaysOfRent(startDate, endDate string) (int64, error) {

	layout := "2006-01-02"

	startDates, err := time.Parse(layout, startDate)
	if err != nil {
		return 0, err
	}
	endDates, err := time.Parse(layout, endDate)
	if err != nil {
		return 0, err
	}

	if startDates.After(endDates) {
		return 0, errors.New("error start date after end date")
	}

	daysOfRent := int64((endDates.Sub(startDates).Hours() / 24) + 1)
	return daysOfRent, nil

}

func calculateDiscount(daysOfRent int64, dailyRent float64, membershipId *int64) float64 {

	var discountRate float64

	var membership models.Membership

	if membershipId == nil {
		membership.Name = ""
	} else {
		if err := config.DB.First(&membership, membershipId).Error; err != nil {
			return 0
		}
	}

	switch membership.Name {
	case "bronze":
		discountRate = 0.04
	case "silver":
		discountRate = 0.07
	case "gold":
		discountRate = 0.15
	default:
		discountRate = 0
	}

	return float64(daysOfRent) * dailyRent * discountRate
}

// Calculate driver incentive
func calculateDriverIncentive(daysOfRent int64, dailyCarRent float64) float64 {
	return float64(daysOfRent) * dailyCarRent * 0.05 // 5% incentive
}
