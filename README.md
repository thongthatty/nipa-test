# Nipa-test

nipa-test is a simple test for CRUD

# Getting Started
These instructions will provide you the information to setup this project on your local machine.

Also, please note that this project's scripts are written for Unix-based OS. So some commands or Makefile will not work on Windows.
We are open for everyone to contribute to improve this project.

## Prerequisites
Please make sure you already have the following software installed.
- [Go 1.15+](https://golang.org/dl/) 

#### Install prerequisite tools
For MacOS use the following
```sh
brew install go
```

#### Running the test
Run `Unit Test` and check code coverage by executing following command.
```sh
make test
```

#### Run dev
```sh
make dev
```

## Regenerate code, then test
```sh
make all
```

## Api doc
http://localhost:1323/swagger/index.html


## Example
```sh
curl --location --request GET 'http://localhost:1323/api/ticket'
curl --location --request GET 'http://localhost:1323/api/ticket?status=PENDING&from=${unix_time}&to=${unix_time}&page=1&page_size=10'
curl --location --request PUT 'http://localhost:1323/api/ticket/1' \
--header 'Content-Type: application/json' \
--data-raw '{
	"status": "REJECTED"
}'
curl --location --request POST 'http://localhost:1323/api/ticket' \
--header 'Content-Type: application/json' \
--data-raw '{
	"name": "test"
}'
```