#!/bin/bash

GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -o app/go-web-server app/go-web-server.go
docker build --no-cache -t go-web-server app/
docker run --name go-ws -p 8000:8080 --rm go-web-server:latest
rm app/go-web-server
