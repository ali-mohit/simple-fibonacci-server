
fibo_app:
	go build -o $@ ./cmd

test: ## to run tests
	go test ./...

# Variables
ROOT := github.com/ali-mohit/simple-fibonacci-server