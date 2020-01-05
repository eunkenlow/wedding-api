package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/eunkenlow/wedding-api/shared/apperror"
	"github.com/eunkenlow/wedding-api/shared/util/validator"
)

// New returns new echo router
func New() *echo.Echo {
	e := echo.New()

	// Middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	e.Use(middleware.RequestID())
	e.Use(middleware.Recover())
	// Validator
	e.Validator = validator.New()
	// Error handling
	e.HTTPErrorHandler = apperror.HTTPErrorHandler

	return e
}
