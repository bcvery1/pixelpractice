#!/usr/bin/bash

go build -ldflags "-s -w" -o dist/server.exe server/main.go
go build -ldflags "-H=windowsgui -s -w" -o dist/client.exe client/main.go
