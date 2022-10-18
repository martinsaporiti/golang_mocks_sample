BIN := $(shell pwd)/bin
GO?=$(shell which go)
export GOBIN := $(BIN)
export PATH := $(BIN):$(PATH)

$(BIN)/mockery: go.mod go.sum ## install code generator for API files.
	$(GO) install github.com/vektra/mockery/v2@latest


generate/mocks:  $(BIN)/mockery
	mockery --all --keeptree