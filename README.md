# Introduction

ERP Prototype BE is a **REST API for the ERP Prototype Web Project**.

# Requirements

This module requires the following toolkit(s):

1. **Golang**: SDK 1.21.6
2. **PostgreSQL**: v17.2

## Installation & Running

1. Clone the repository ```git clone https://github.com/TheRealMarcell/erp-proto-be```
2. Run `go mod tidy` or by comprehensive way ```go mod download```
3. Create an .env file on the root directory to store database configurations
4. To start the application, run ```go run cmd/api.go``` or ```make run```
5. Once run, to view the swagger documentation, navigate to ```http://localhost:8080/swagger/index.html```

## Running with Air (Dev build)
1. Initialise with ```air init```
2. Configure the .toml file by modifying the cmd to ```go build -o ./tmp/main ./cmd```
3. Run with live-reload using ```air -c .air.toml``` or simply ```air```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.