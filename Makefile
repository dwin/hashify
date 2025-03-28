# Makefile for building and testing the project

# Variables
CARGO = cargo

# Targets
.PHONY: all build test clean

all: build

build:
	$(CARGO) build --release

test:
	$(CARGO) test

clean:
	$(CARGO) clean
