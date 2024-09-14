clean:
	@echo "Cleaning up..."
	@rm -rf ./main

build:
	@echo "Building the application..."
	@go build cmd/api/main.go

start:
	@echo "Starting the application..."
	@./main

run:
	@echo "Starting the application..."
	@go run cmd/api/main.go

swagger:
	@echo "Generating swagger.yaml"
	@swag init -g cmd/api/main.go -o ./docs


startApp: clean swagger build start
runApp: swagger run