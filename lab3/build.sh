#!/bin/bash
go build ./src/cmd/main.go
mkdir -p ./bin
mv ./main ./bin/peer
