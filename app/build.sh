#!/bin/bash

GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -o go-web-server go-web-server.go || exit 1
docker build --no-cache -t aracki/go-web-server . || exit 1
docker push aracki/go-web-server
