#!/usr/bin/env bash

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./linux_amd64 ./Dujt.go
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ./darwin_amd64 ./Dujt.go
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ./windows_amd64.exe ./Dujt.go

CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=5 go build -o ./linux_armv5 ./Dujt.go
CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=6 go build -o ./linux_armv6 ./Dujt.go
CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=7 go build -o ./linux_armv7 ./Dujt.go
