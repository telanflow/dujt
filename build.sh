#!/usr/bin/env bash

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./Dujt_linux_amd64 ./Dujt.go
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ./Dujt_darwin_amd64 ./Dujt.go
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ./Dujt_windows_amd64.exe ./Dujt.go
