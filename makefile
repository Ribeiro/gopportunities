.PHONY: default run build test docs clean
# Variables
APP_NAME=gopportunities

# Tasks
default: run-with-docs

run:
	@go run cmd/api/main.go
run-with-docs:
	@swag init
	@go run cmd/api/main.go
build:
	@go build -o $(APP_NAME) cmd/api/main.go
test:
	@go test ./ ...
docs:
	@swag init
clean:
	@rm -f $(APP_NAME)
	@rm -rf cmd/api/docs