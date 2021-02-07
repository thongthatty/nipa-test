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
