package main

import (
	"github.com/go-pg/migrations/v7"
	"go.uber.org/zap"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		zap.S().Info("creating table rsvps...")
		_, err := db.Exec(`
			CREATE TABLE rsvps(
				id bigserial,
				first_name VARCHAR(255) NOT NULL,
				last_name VARCHAR(255) NOT NULL,
				email VARCHAR(255) NOT NULL,
				phone_number VARCHAR(24) NOT NULL,
				attending BOOLEAN NOT NULL DEFAULT false,
				notifications BOOLEAN NOT NULL DEFAULT false,
				dietary_requirements VARCHAR(255),
				best_moments TEXT,

				created_at timestamptz NOT NULL DEFAULT now(),

				PRIMARY KEY (id)
			);
		`)
		return err
	}, func(db migrations.DB) error {
		zap.S().Info("dropping table rsvps...")
		_, err := db.Exec(`DROP TABLE rsvps`)
		return err
	})
}
