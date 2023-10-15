ifeq ($(POSTGRES_SETUP_TEST),)
	POSTGRES_SETUP_TEST := user=test password=test dbname=test host=localhost port=5432 sslmode=disable
endif

INTERNAL_PATH=$(CURDIR)/internal
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

.PHONY: up
up:
	docker-compose up

generate:
	rm -rf internal/proto
	mkdir -p internal/proto

	protoc \
		--proto_path=internal/grpc_api/proto \
		--go_out=internal/proto \
		--go-grpc_out=internal/proto \
		internal/grpc_api/proto/*.proto

