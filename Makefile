include .env-file
export

tidy:
	go mod tidy

run:
	go run cmd/tvseries-cli/main.go

run-race:
	go run -race cmd/tvseries-cli/main.go

test:
	go test -cover -race ./...