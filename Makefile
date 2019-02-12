GO := go
PKGS := $(shell $(GO) list ./...)
SERVER_NAME := server
WORKER_NAME := worker
BUILD_DIR ?= build

all: build test

build-server:
	@echo ">> building $(SERVER_NAME) binaries"
	$(GO) build -o $(BUILD_DIR)/$(SERVER_NAME) ./$(SERVER_NAME)

build-worker:
	@echo ">> building $(WORKER_NAME) binaries"
	$(GO) build -o $(BUILD_DIR)/$(WORKER_NAME) ./$(WORKER_NAME)

build: build-server build-worker

test:
	@echo ">> testing binaries"
	$(GO) test -short -race $(PKGS)

clean:
	@echo ">> removing binaries"
	rm -rf $(BUILD_DIR)

.PHONY: build-server build-worker build test clean
