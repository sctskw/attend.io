#! /usr/bin/make
#
# Makefile for attend.io

.PHONY: docs build test clean

NAME=attend.io

build:
	@swagger generate server -A $(NAME) -f ./swagger.yml

docs:
	@npx widdershins ./swagger.yml -o docs.md --language_tabs "shell:Shell"
	@swagger-codegen generate -i ./swagger.yml -l html2 -o docs

serve:
	@go run ./cmd/attend-io-server/main.go

test: export GOOGLE_APPLICATION_CREDENTIALS=./creds.json
	@gotestsum --format short-verbose ./tests/...

test-watch: export GOOGLE_APPLICATION_CREDENTIALS=./creds.json
	@gotestsum --format short-verbose --watch ./tests/...

clean:
	-rm docs.md

