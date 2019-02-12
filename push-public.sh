#!/bin/bash

GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -o app/go-web-server app/go-web-server.go
docker build --no-cache -t aracki/go-web-server app/
rm app/go-web-server
docker push aracki/go-web-server
