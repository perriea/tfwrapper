########################################
#          TFVERSION MAKEFILE          #
#        Author: Aurelien PERRIER      #
########################################

GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
BIN=tfwrapper

DOCKERCMD=docker
DOCKERBUILD=$(DOCKERCMD) build
CONTNAME=perriea/tfwrapper:latest

.PHONY: help test build docker vendor-list vendor-update

all: build

docker:
	@echo "Build binary & Docker image"
	@GOOS=linux $(GOBUILD) -i -o ./$(BIN) ./
	$(DOCKERBUILD) . -t $(CONTNAME)

test:
	$(GOTEST) $(go list ./... | grep -v '/vendor/')

build:
	@echo "Build binary"
	@$(GOBUILD) -i -o ./$(BIN) ./

vendor-list:
	@govendor list

vendor-update:
	@govendor update +vendor