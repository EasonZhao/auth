# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOMOD=$(GOCMD) mod
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOLINT=$(GOCMD)lint
BINARY_NAME=mybinary
BINARY_UNIX=$(BINARY_NAME)_unix

.PHONY: all build help

build:
	swag init
	$(GOBUILD) -v .

help:
	@echo "make: compile packages and dependencies"
	@echo "make clean: remove object files and cached files"