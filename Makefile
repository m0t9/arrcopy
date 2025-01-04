format:
	gofumpt -w .

unit-test:
	go test -race ./...
