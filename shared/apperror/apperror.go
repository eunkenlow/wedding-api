package apperror

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// AppError - Error details struct
type AppError struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	Code    int         `json:"code"`              // application-specific error code
	Message interface{} `json:"message"`           // application-level error message
	Details []string    `json:"details,omitempty"` // application-level details message
}

// Error - Cast type to error
func (e *AppError) Error() string {
	return e.Message.(string)
}

func new(code int, message string, httpStatus int) *AppError {
	return &AppError{
		Err:            errors.New(message),
		HTTPStatusCode: httpStatus,
		Code:           code,
		Message:        message,
	}
}

// HTTPErrorHandler - custom http error handler
func HTTPErrorHandler(err error, c echo.Context) {
	// Echo errors
	if he, ok := err.(*echo.HTTPError); ok {
		c.JSON(he.Code, he)
		return
	}
	// Undefined errors
	apperror := ErrInternalServerError
	// App errors
	if err, ok := err.(*AppError); ok {
		apperror = err
	}
	// Validation errors
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		err := ErrInvalidRequest
		err.Details = make([]string, 0)
		for _, ve := range validationErrors {
			field := ve.Field()
			tag := ve.Tag()
			if tag == "required" {
				err.Details = append(err.Details, fmt.Sprintf("%s is %s", field, tag))
			} else {
				err.Details = append(err.Details, fmt.Sprintf("%s failed on '%s' tag", field, tag))
			}
		}
		apperror = err
	}
	// Return error
	c.JSON(apperror.HTTPStatusCode, apperror)
}

// 1XXXX - General errors
var (
	ErrBadRequest          = new(10001, "Bad request", http.StatusBadRequest)
	ErrUnauthorized        = new(10002, "Unauthorized", http.StatusUnauthorized)
	ErrInternalServerError = new(10003, "Something went wrong", http.StatusInternalServerError)
	ErrInvalidRequest      = new(10004, "Invalid request", http.StatusBadRequest)
)
