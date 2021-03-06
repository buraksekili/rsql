#!/bin/sh
cd ./cmd || exit
go build -o $GOPATH/bin/rsql
