MODE ?= DEV

export MODE

build:
	@echo "Building in $(MODE) mode"
	go build -tags=$(MODE)
	
clean:
	go clean

test:
	go test ./...