do_all:
	./test.sh
	GOOS=linux go build -o bin/subscription cmd/subscription/main.go
	GOOS=linux go build -o bin/webhook cmd/webhook/main.go
	GOOS=linux go build -o bin/activity cmd/activity/main.go
	GOOS=linux go build -o bin/oauth cmd/oauth/main.go
.PHONY: do_all