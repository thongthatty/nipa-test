GOCMD=go
GORUN=$(GOCMD) run
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
DIST=./dist/ticketsupport 
all: clean test build
build: 
		$(GOBUILD) -o ${DIST} ./cmd/ticketsupport/main.go
test: 
		$(GOTEST) -v ./internal/*
clean: 
		$(GOCLEAN)
		rm -f $(BINARY_NAME)
run:
		$(GOBUILD) -o $(BINARY_NAME) -v ./
		./$(BINARY_NAME)
dev:	
		$(GORUN) cmd/ticketsupport/main.go

		