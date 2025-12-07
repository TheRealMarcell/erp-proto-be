# Introduction

ERP Prototype BE is a **REST API for the ERP Prototype Web Project**.

# Requirements

This module requires the following toolkit(s):

1. **Golang**: SDK 1.21.6
2. **PostgreSQL**: v17.2
3. 

## Installation & Running

1. Clone the repository ```git clone https://github.com/TheRealMarcell/erp-proto-be```
2. Run `go mod tidy` or by comprehensive way ```go mod download```
3. Create an .env file on the root directory to store database configurations
4. To start the application, run ```go run cmd/api.go``` or ```make run```
5. To generate swagger documentation, run ```swag init -g cmd/api.go -o docs```
6. Once run, to view the swagger documentation, navigate to ```http://localhost:8080/swagger/index.html```
7. To view the application's web interface, navigate to ```http://localhost:8080/web```
8. Provide a dummy login (username: user, password: user)

## Configuring the PostgreSQL database locally
1. Ensure you have installed and initialized a PostgreSQL database on your local machine
2. Have a look at the env.example file, and create a .env file in the root directory with your own permission configurations

## Using docker to run service
1. Ensure you have installed docker on your machine
2. Launch the docker app
3. Start the docker service for your postgreSQL instance: ```docker start database-container```
4. To build the service or run to reflect changes, ```make docker-up``` or ```docker compose up --build```
5. To tear down and flush the database, run ```docker compose down -v```
6. To simply run the code (no changes rendered), run ```docker compose up```
7. To simply close the containers, run ```docker compose down```

## Running with Air (Dev build)
1. Initialise with ```air init```
2. Configure the .toml file by modifying the cmd to ```go build -o ./tmp/main ./cmd```
3. Run with live-reload using ```air -c .air.toml``` or simply ```air```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.