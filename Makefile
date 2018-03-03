########################################
#          TFWRAPPER MAKEFILE          #
#        Author: Aurelien PERRIER      #
########################################

GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
BIN=tfwrapper

DOCKERCMD=docker
DOCKERBUILD=$(DOCKERCMD) build
CONTNAME=perriea/tfwrapper:latest

all: build

docker:
	@echo "Build binary ..."
	@GOOS=linux $(GOBUILD) -i -o ./$(BIN) ./
	@echo "Build Docker image ..."
	$(DOCKERBUILD) . -t $(CONTNAME)

test:
	@echo "Testing ..."
	$(GOTEST) `go list ./... | grep -v '/vendor/'`

build:
	@echo "Build binary"
	@$(GOBUILD) -i -o ./$(BIN) ./

vendor-list:
	@govendor list

vendor-update:
	@govendor update +vendor

.PHONY: help test build docker vendor-list vendor-update