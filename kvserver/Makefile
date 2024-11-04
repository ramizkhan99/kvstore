# Golang Makefile for kvserver which accepts port as an argument

port = 50051

.PHONY: tidy

run:
	go run src/main.go -port $(port)

tidy:
	go mod tidy

build:
	go build src/main.go