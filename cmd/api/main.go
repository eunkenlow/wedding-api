package main

import (
	"github.com/eunkenlow/wedding-api/config"
	"github.com/eunkenlow/wedding-api/service/api"
	"github.com/eunkenlow/wedding-api/shared/util/logger"
)

// Api microservice server
func main() {
	logger := logger.New()
	defer logger.Sync()
	cfg := config.New()
	api.Start(cfg)
}
