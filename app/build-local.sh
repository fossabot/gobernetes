#!/bin/bash

GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -o go-web-server go-web-server.go || exit 1
docker build --no-cache -t go-web-server . || exit 1
docker run --name go-ws -p 8000:8080 --rm go-web-server:latest
