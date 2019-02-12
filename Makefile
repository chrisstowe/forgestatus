GO := go
SERVER_NAME := server
WORKER_NAME := worker

# Set this var to override the build directory.
BUILD_DIR ?= build

all: test install

get-dep-server:
	@echo ">> getting $(SERVER_NAME) dependencies"
	$(GO) get -d ./$(SERVER_NAME)

get-dep-worker:
	@echo ">> getting $(WORKER_NAME) dependencies"
	$(GO) get -d ./$(WORKER_NAME)

build-server: get-dep-server
	@echo ">> building $(SERVER_NAME) binaries"
	$(GO) build -o $(BUILD_DIR)/$(SERVER_NAME) ./$(SERVER_NAME)

build-worker: get-dep-worker
	@echo ">> building $(WORKER_NAME) binaries"
	$(GO) build -o $(BUILD_DIR)/$(WORKER_NAME) ./$(WORKER_NAME)

build: build-server build-worker

install: get-dep-server get-dep-worker
	@echo ">> installing binaries"
	$(GO) install ./server ./worker

test: get-dep-server get-dep-worker
	@echo ">> testing binaries"
	$(GO) test -v ./common ./server ./worker

clean:
	@echo ">> removing binaries"
	rm -rf $(BUILD_DIR)

.PHONY: get-dep-server get-dep-worker build-server build-worker build install test clean
