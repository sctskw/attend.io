#! /usr/bin/make
#
# Makefile for attend.io


.PHONY: format build compile

NAME=attend.io

build:
	@swagger generate server -A $(NAME) -f ./swagger.yml

docs:
	@npx swagger-markdown -i ./swagger.yml -o docs.md
	@swagger-codegen generate -i ./swagger.yml -l html2 -o docs

serve:
	@go run ./cmd/attend-io-server/main.go

clean:

