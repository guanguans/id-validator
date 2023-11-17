# note: call scripts from /scripts
GOPATH=/opt/homebrew/bin/go=go

golangci-lint:
	golangci-lint run ./... --color=always --verbose

gosec:
	gosec ./... -color -verbose

fmt:
	go fmt ./...

vet:
	go vet ./..。

test:
	go test ./... -cover -coverprofile=cover.out -v

bench:
	go test ./... -bench=. -benchmem -v

cover:
	go tool cover -html=cover.out
