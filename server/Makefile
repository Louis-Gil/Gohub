MODE ?= DEV

export MODE

build:
	@echo "Building in $(MODE) mode"
	go build -tags=$(MODE)

build-prod: MODE=PROD
build-prod:
	@echo "Building in $(MODE) mode"
	GOOS=linux GOARCH=amd64 go build -tags=$(MODE)
	
clean:
	go clean

test:
	go test .