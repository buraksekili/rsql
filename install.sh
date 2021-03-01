#!/bin/sh
cd ./cmd/cli || exit
go build -o $GOPATH/bin/rsql