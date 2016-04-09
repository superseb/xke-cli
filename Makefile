.PHONY: cover

default: test
	go build

test:
	go test ./...

cover:
	go test ./xke -coverprofile=cover.out
	go tool cover -html=cover.out
	rm cover.out

release: osx linux windows

osx:
	GOOS=darwin GOARCH=amd64 go build
	mv xke-cli xke-cli-darwin

linux:
	GOOS=linux GOARCH=amd64 go build
	mv xke-cli xke-cli-linux

windows:
	GOOS=windows GOARCH=amd64 go build
