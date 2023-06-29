# Makefile for the docscan Go application

# Phony targets ensure that make doesn't confuse these targets with files of the same name
.PHONY: build clean run mod

# Variables
APP_NAME=docscan
GO_FILES=$(wildcard *.go)

# The default target, if no target is specified to make
default: build

# The mod target is responsible for setting up the Go module if it doesn't already exist
# If it doesn't, initialize a new Go module and get the necessary dependencies
mod:
	# Check if the go.mod file exists
	if [ ! -f go.mod ]; then \
		go mod init your_module_name; \
		go get github.com/fatih/color; \
	fi

# The build target builds the Go application
build: mod
	go build -o $(APP_NAME) $(GO_FILES)

# The run target runs the built Go application
run: build
	./$(APP_NAME)

# The clean target removes the built Go application
# Check if the application file exists
# If it does, remove it
clean:
	if [ -f $(APP_NAME) ]; then \
		rm $(APP_NAME); \
	fi
