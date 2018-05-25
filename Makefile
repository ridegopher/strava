do_all:
	go test -v ./...
	GOOS=linux go build -o bin/subscription cmd/subscription/main.go
	GOOS=linux go build -o bin/webhook cmd/subscription/main.go
.PHONY: do_all