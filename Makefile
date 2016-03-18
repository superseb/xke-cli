.PHONY: cover

test:
	go test ./...

cover:
	go test ./xke -coverprofile=cover.out
	go tool cover -html=cover.out
	rm cover.out
