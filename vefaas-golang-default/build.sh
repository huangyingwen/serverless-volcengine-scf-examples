#! /bin/bash
set -ex

GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main
