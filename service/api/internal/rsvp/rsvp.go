package rsvp

import (
	"github.com/go-pg/pg/v9"
	"github.com/labstack/echo/v4"
)

// New create rsvp service routes.
func New(e *echo.Group, db *pg.DB) {
	repo := NewRepository(db)
	svc := NewService(repo)
	h := NewHandler(svc)

	e.POST("/rsvps", h.PostRsvps())
	e.GET("/rsvps", h.GetRsvps())
}
