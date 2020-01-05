package db

import (
	"fmt"

	"github.com/eunkenlow/wedding-api/config"
	"github.com/go-pg/pg/v9"
	"go.uber.org/zap"
)

// Connect - Initialize a new connection to db
func Connect(cfg *config.Config) *pg.DB {
	opts := &pg.Options{
		User:     cfg.DBUser,
		Password: cfg.DBPass,
		Database: cfg.DBName,
		Addr:     fmt.Sprintf("%s:%d", cfg.DBHost, cfg.DBPort),
	}
	db := pg.Connect(opts)
	err := checkConn(db)
	if err != nil {
		zap.S().Fatalf("Error connecting to %s db", cfg.DBName)
		db.Close()
	}
	zap.S().Infof("Successfully connected to %s db\n", cfg.DBName)
	return db
}

func checkConn(db *pg.DB) error {
	var n int
	_, err := db.QueryOne(pg.Scan(&n), "SELECT 1")
	return err
}
