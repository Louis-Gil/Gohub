MODE ?= DEV

export MODE

build:
	@echo "Building in $(MODE) mode"
	go build -tags=$(MODE)

build-prod: MODE=PROD
build-prod:
	@echo "Building in $(MODE) mode"
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags="$(MODE) lambda.norpc" -o server main.go 
	
build-test: MODE=TEST
build-test:
	@echo "Building in $(MODE) mode"
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags="$(MODE) lambda.norpc" -o server main.go 

clean:
	go clean

test:
	go test .