install: 
	@echo "Downloading all dependencies/libraries"
	go mod download

run:
	@echo "Running the application"
	go run cmd/api.go

tidy: 
	@echo "Cleaning and updating go.mod"
	go mod tidy

swag-docs:
	@echo "Generating swagger documentation"
	swag init -g cmd/api.go