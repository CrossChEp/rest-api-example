POSTGRES_SETUP_TEST := user=postgres password=test database=tips_db host=localhost port=54321 sslmode=disable

MIGRATION_FOLDER=$(CURDIR)/migrations

.PHONY: migration-create
migration-create:
	goose -dir "$(MIGRATION_FOLDER)" create "$(name)" sql

.PHONY: test-migration-up
test-migration-up:
	goose -dir "$(MIGRATION_FOLDER)" postgres "$(POSTGRES_SETUP_TEST)" up

.PHONY: test-migration-down
test-migration-down:
	goose -dir "$(MIGRATION_FOLDER)" postgres "$(POSTGRES_SETUP_TEST)" down

.PHONY: linter-check
linter-check:
	golangci-lint run
