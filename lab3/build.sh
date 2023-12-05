#!/bin/bash
go mod tidy
go build ./src/cmd/main.go
mkdir -p ./bin
mv ./main ./bin/peer
