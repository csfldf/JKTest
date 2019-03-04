GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=$(shell pwd | awk -F '/' '{print $$(NF-1)}')

all:
	@echo $(GOPATH)
	$(GOBUILD) -o ../bin/$(BINARY_NAME) -v ./launcher/main.go 
test:
	@echo $(BINARY_NAME)