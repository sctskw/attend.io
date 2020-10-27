#! /usr/bin/make
#
# Makefile for attend.io


.PHONY: format build compile


NAME=attend.io

build:
	@swagger generate server -A $(NAME) -f ./swagger.yml


serve:
	@go run ./cmd/attend-io-server/main.go
