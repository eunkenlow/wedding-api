package main

import (
	"flag"

	"github.com/eunkenlow/wedding-api/config"
	"github.com/eunkenlow/wedding-api/shared/db"
	"github.com/eunkenlow/wedding-api/shared/util/logger"
	"github.com/go-pg/migrations/v7"
	"go.uber.org/zap"
)

func main() {
	flag.Parse()
	logger := logger.New()
	defer logger.Sync()
	cfg := config.New()
	db := db.Connect(cfg)

	// Init go-pg migrations if version is 0
	version, err := migrations.Version(db)
	if version == 0 {
		migrations.Run(db, "init")
	}

	oldVersion, newVersion, err := migrations.Run(db, flag.Args()...)
	if err != nil {
		zap.S().Fatal(err.Error())
	}
	if newVersion != oldVersion {
		zap.S().Infof("migrated from version %d to %d\n", oldVersion, newVersion)
	} else {
		zap.S().Infof("version is %d\n", oldVersion)
	}
}
