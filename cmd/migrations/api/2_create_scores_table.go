package main

import (
	"github.com/go-pg/migrations/v7"
	"go.uber.org/zap"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		zap.S().Info("creating table scores...")
		_, err := db.Exec(`
			CREATE TABLE scores(
				id bigserial,
				name VARCHAR(255) NOT NULL,
				score INT NOT NULL,

				PRIMARY KEY (id)
			);
		`)
		return err
	}, func(db migrations.DB) error {
		zap.S().Info("dropping table scores...")
		_, err := db.Exec(`DROP TABLE scores`)
		return err
	})
}
