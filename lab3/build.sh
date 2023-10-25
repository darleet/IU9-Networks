#!/bin/bash
cd ./src
go build
mkdir -p ../bin
mv ./peer ../bin/peer
