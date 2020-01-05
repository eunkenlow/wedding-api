package score

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

// PostScore creates new high score
func (h *Handler) PostScore() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := &Score{}
		if err := c.Bind(req); err != nil {
			return apperror.ErrInvalidRequest
		}
		if err := c.Validate(req); err != nil {
			return err
		}

		if err := h.service.CreateScore(req); err != nil {
			return err
		}

		return c.NoContent(http.StatusNoContent)
	}
}

// GetScores returns top 10 high scores
func (h *Handler) GetScores() echo.HandlerFunc {
	return func(c echo.Context) error {
		scores, err := h.service.FindAllScores()
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, scores)
	}
}
