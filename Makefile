run:
	go run cmd/api/main.go

pretty:
	gofmt -s -w .

test:
	go test -race ./...
