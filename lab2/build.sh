#!/bin/bash
cd ./src
go get github.com/mgutz/logxi/v1
go get golang.org/x/net/html
go install .

go build
mv ./server ../bin/server