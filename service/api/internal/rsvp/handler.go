package rsvp

import (
	"net/http"

	"github.com/eunkenlow/wedding-api/shared/apperror"
	"github.com/labstack/echo/v4"
)

// Handler represents the handler for rsvp service.
type Handler struct {
	service Service
}

// NewHandler initialises a new handler for rsvp service.
func NewHandler(service Service) *Handler {
	return &Handler{service}
}

// PostRsvps creates a new rsvp
func (h *Handler) PostRsvps() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := &Rsvp{}
		if err := c.Bind(req); err != nil {
			return apperror.ErrInvalidRequest
		}
		if err := c.Validate(req); err != nil {
			return err
		}

		if err := h.service.CreateRsvp(req); err != nil {
			return err
		}

		return c.NoContent(http.StatusNoContent)
	}
}

// GetRsvps returns all rsvps
func (h *Handler) GetRsvps() echo.HandlerFunc {
	return func(c echo.Context) error {
		rsvps, err := h.service.FindAllRsvps()
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, rsvps)
	}
}
