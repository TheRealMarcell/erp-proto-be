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

dev:
	@echo "Running dev mode with air live-reload"
	air 

docker-up:
	@echo "Building the docker image, running database and server"
	docker compose up -d
	sleep 1
	open http://localhost:8080/web