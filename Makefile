format:
	gofumpt -w .

unit-test:
	go test -race ./...

coverage:
	go test ./... -coverprofile=./cover.out -covermode=atomic -coverpkg=./...
	go tool cover -func=cover.out

