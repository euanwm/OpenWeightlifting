# Backend Makefile.

.DEFAULT_GOAL := run

# Compile and launch the backend server in local development mode.
.PHONY: run
run:
	go run backend local

# Execute backend unit tests.
.PHONY: test
test:
	go test ./...

# Regenerate Swaggo docs from endpoint annotations.
.PHONY: docs
docs:
	swag init --parseDependency --parseInternal -g main.go
