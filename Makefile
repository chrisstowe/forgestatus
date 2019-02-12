GO := go
SERVER_NAME := server
WORKER_NAME := worker

# Set this var to override the build directory.
BUILD_DIR ?= build

all: test install

get-dep:
	@echo ">> getting dependencies"
	$(GO) get -d ./$(SERVER_NAME)
	$(GO) get -d ./$(WORKER_NAME)

build-server: get-dep
	@echo ">> building $(SERVER_NAME) binaries"
	$(GO) build -o $(BUILD_DIR)/$(SERVER_NAME) ./$(SERVER_NAME)

build-worker: get-dep
	@echo ">> building $(WORKER_NAME) binaries"
	$(GO) build -o $(BUILD_DIR)/$(WORKER_NAME) ./$(WORKER_NAME)

build: build-server build-worker

install: get-dep
	@echo ">> installing binaries"
	$(GO) install ./server ./worker

test: get-dep
	@echo ">> testing binaries"
	$(GO) test -v ./common ./server ./worker

clean:
	@echo ">> removing binaries"
	rm -rf $(BUILD_DIR)

.PHONY: get-dep build-server build-worker build install test clean
