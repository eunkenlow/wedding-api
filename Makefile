-include .env
export

# Go parameters
GOCMD=go
GORUN=$(GOCMD) run
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test

start:
	$(GORUN) cmd/api/main.go

migrate-api:
	$(GORUN) cmd/migrations/api/*.go

rollback-api:
	$(GORUN) cmd/migrations/api/*.go down
