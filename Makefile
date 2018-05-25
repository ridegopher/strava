do_all:
	./test.sh
	GOOS=linux go build -o bin/subscription cmd/subscription/main.go
	GOOS=linux go build -o bin/webhook cmd/webhook/main.go
.PHONY: do_all