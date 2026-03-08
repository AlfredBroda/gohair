ifeq (, $(shell go env GOPATH))
BINDIR = ./bin
else
BINDIR = $(shell go env GOPATH)/bin
endif

-include .env
export

.PHONY: all
all: lint build

.PHONY: build
build:
	go build -o bin/server ./cmd/server/main.go

.PHONY: docker-build
docker: lint
	docker compose -f docker-compose.yml up -d

.PHONY: run
run: build
	./bin/server

.PHONY: dev
dev: lint
	air

.PHONY: migrate-up
migrate-up:
	goose -dir $(GOOSE_MIGRATION_DIR) $(GOOSE_DRIVER) $(GOOSE_DBSTRING) up

.PHONY: migrate-down
migrate-down:
	goose -dir $(GOOSE_MIGRATION_DIR) $(GOOSE_DRIVER) $(GOOSE_DBSTRING) down

.PHONY: migrate-create
migrate-create:
	goose -s -dir $(GOOSE_MIGRATION_DIR) create migration sql

.PHONY: lint
lint:
ifeq (, $(shell $(BINDIR)/golangci-lint --version))
	@echo "golangci-lint not found, installing..."
	wget -O- -nv https://raw.githubusercontent.com/golangci/golangci-lint/HEAD/install.sh | sh -s -- -b $(BINDIR) v2.7.2
endif
	$(BINDIR)/golangci-lint run ./...

.PHONY: clean
clean:
	rm -rf ./bin
