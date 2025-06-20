APP_NAME=risk-detection
CMD_DIR=./cmd/server
BIN_DIR=./bin

.PHONY: build run clean test

build:
	@echo "==> Building $(APP_NAME)..."
	@mkdir -p $(BIN_DIR)
	@go build -o $(BIN_DIR)/$(APP_NAME) $(CMD_DIR)

run:
	@echo "==> Running $(APP_NAME)..."
	@go run $(CMD_DIR)

clean:
	@echo "==> Cleaning..."
	@rm -rf $(BIN_DIR)
