# WIRE
# Define the Go binary (you might want to override this for cross-compilation)
GO ?= go

# Define the Wire binary
WIRE ?= wire

# Define output binary name
BINARY_NAME := example-app

# Define your Go main package
MAIN_PACKAGE := ./cmd

.PHONY: all wire build clean

#all: wire build

# Run wire to generate dependency injection code
wire:
	@echo "Running Wire..."
	$(WIRE) ./injector

# Build the Go application
build: wire
	@echo "Building the application..."
	$(GO) build -o $(BINARY_NAME) $(MAIN_PACKAGE)

# Clean up generated files
clean:
	@echo "Cleaning up..."
	$(GO) clean
	rm -f $(BINARY_NAME)
	rm -f wire_gen.go
# END WIRE