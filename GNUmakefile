########################################
#          TFWRAPPER MAKEFILE          #
#        Author: Aurelien PERRIER      #
########################################

GOCMD=go
DOCKERCMD=docker

SOURCEDIR=.
SOURCES := $(shell find $(SOURCEDIR) -name '*.go' | grep -v vendor)
BIN_FOLDER=bin/
BIN=$(BIN_FOLDER)tfwrapper

GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test

DOCKERBUILD=$(DOCKERCMD) build
CONTNAME=perriea/tfwrapper:latest

all: build

build: $(SOURCES)
	@echo "Build binary"
	@$(GOBUILD) -i -o ./$(BIN) ./

install: fmt
	go install

fmt:
	gofmt -w $(SOURCES)

vendor-status:
	@govendor status

vendor-list:
	@govendor list

vendor-update:
	@govendor update +vendor

docker:
	@echo "Build binary ..."
	@GOOS=linux $(GOBUILD) -i -o ./$(BIN) ./
	@echo "Build Docker image ..."
	$(DOCKERBUILD) . -t $(CONTNAME)

test:
	@echo "Testing ..."
	$(GOTEST) `go list ./... | grep -v '/vendor/'`

vet:
	@echo "go vet ."
	@go vet $$(go list ./... | grep -v vendor/) ; if [ $$? -eq 1 ]; then \
		echo ""; \
		echo "Vet found suspicious constructs. Please check the reported constructs"; \
		echo "and fix them if necessary before submitting the code for review."; \
		exit 1; \
	fi

clean:
	$(RM) ${BIN}

.PHONY: help build install fmt vendor-status vendor-list vendor-update docker test vet clean