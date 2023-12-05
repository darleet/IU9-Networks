#!/bin/bash
go mod tidy

go build ./src/server.go
mkdir -p ./bin
mv ./server ./bin/server
