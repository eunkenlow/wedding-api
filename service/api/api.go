package api

import (
	"github.com/eunkenlow/wedding-api/config"
	"github.com/eunkenlow/wedding-api/service/api/internal/rsvp"
	"github.com/eunkenlow/wedding-api/service/api/internal/score"
	"github.com/eunkenlow/wedding-api/shared/db"
	"github.com/eunkenlow/wedding-api/shared/router"
)

// Start - Starts running api microservice
func Start(cfg *config.Config) {
	db := db.Connect(cfg)
	r := router.New()

	v1 := r.Group("/api/v1")
	rsvp.New(v1, db)
	score.New(v1, db)

	r.Logger.Fatal(r.Start(":1234"))
}
