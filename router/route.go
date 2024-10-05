package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/reusmana/car-rental-go/controllers"
)

func SetupRoutes(e *echo.Echo) {

	api := e.Group("/api/v1")

	api.GET("/cars", controllers.GetCars)
	api.GET("/cars/:id", controllers.GetCarById)
	api.POST("/cars", controllers.CreateCar)
	api.PUT("/cars/:id", controllers.UpdateCar)
	api.DELETE("/cars/:id", controllers.DeleteCar)

	api.GET("/customer", controllers.GetCustomer)
	api.GET("/customer/:id", controllers.GetCustomerById)
	api.POST("/customer", controllers.CreateCustomer)
	api.PUT("/customer/:id", controllers.UpdateCustomer)
	api.DELETE("/customer/:id", controllers.DeleteCustomer)

	api.GET("/booking", controllers.GetBookings)
	api.GET("/booking/:id", controllers.GetBookingById)
	api.POST("/booking", controllers.CreateBooking)
	api.PUT("/booking/:id", controllers.UpdateBooking)
	api.DELETE("/booking/:id", controllers.DeleteBooking)

	api.GET("/driver", controllers.GetDriver)
	api.POST("/driver", controllers.CreateDriver)
}
