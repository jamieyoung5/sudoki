# Set the name for your output binary
BINARY_NAME=sudoki

# Set the path to your main.go file
MAIN_PACKAGE=./cmd

# Default target, executed when you just type 'make'
all: build

# Builds the binary
build:
	@echo "Building $(BINARY_NAME)..."
	@go build -o $(BINARY_NAME) $(MAIN_PACKAGE)
	@echo "$(BINARY_NAME) built successfully."

# Builds and runs the binary in the TTY
run: build
	@echo "Running $(BINARY_NAME)..."
	@./$(BINARY_NAME)

# Removes the built binary
clean:
	@echo "Cleaning up..."
	@rm -f $(BINARY_NAME)

# Declare targets that are not files
.PHONY: all build run clean