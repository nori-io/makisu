# Makisu Makefile

MAKISU_BUILD_CMD = build -o bin/makisu ./cmd/makisu.go
MAKISU_INSTALL_CMD = install ./cmd/makisu.go

ifeq ($(GO111MODULE),auto)
override GO111MODULE = on
endif

clean: ## remove generated files, tidy vendor dependencies
	@export GO111MODULE=on ;\
	go mod tidy ;\
	rm -rf profile.out ;\
	rm -rf ./bin/makisu
.PHONY: clean

test: ## run tests
	@go test -v ./...
.PHONY: test

build: ## build makisu binary in /bin/makisu
	@go $(MAKISU_BUILD_CMD)
.PHONY: build

install: ## install makisu binary
	@go $(MAKISU_INSTALL_CMD)
.PHONY: install

lint: ## execute linter
	golangci-lint run --no-config --issues-exit-code=0 --deadline=30m \
	  --disable-all --enable=deadcode  --enable=gocyclo --enable=golint --enable=varcheck \
	  --enable=structcheck --enable=maligned --enable=errcheck --enable=dupl --enable=ineffassign \
	  --enable=interfacer --enable=unconvert --enable=goconst --enable=gosec --enable=megacheck ./... ;
.PHONY: lint

help: ## run this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
.PHONY: help

.DEFAULT_GOAL := help