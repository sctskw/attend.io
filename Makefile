#! /usr/bin/make
#
# Makefile for attend.io


.PHONY: format build compile

NAME=attend.io

build:
	@swagger generate server -A $(NAME) -f ./swagger.yml

docs:
	@swagger-codegen generate -i ./swagger.yml -l html2 -o docs

serve:
	@go run ./cmd/attend-io-server/main.go

clean:

