package utils

import "github.com/labstack/echo/v4"

// APIResponse is a standardized format for responses
type APIResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"` // `omitempty` excludes field if it's empty
}

// JSONResponse returns a JSON response with status, message, and optional data
func JSONResponse(c echo.Context, status int, message string, data interface{}) error {
	response := APIResponse{
		Status:  status,
		Message: message,
		Data:    data,
	}
	return c.JSON(status, response)
}
