ifeq (, $(shell go env GOPATH))
BINDIR = ./bin
else
BINDIR = $(shell go env GOPATH)/bin
endif

.PHONY: all
all: lint build

.PHONY: build
build:
	go build -o bin/server ./main.go

.PHONY: docker-build
docker-build: lint
	docker build .

.PHONY: run
run: build
	./bin/server

.PHONY: dev
dev: lint
	air

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
