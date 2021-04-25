# note: call scripts from /scripts
GOCMD=GO111MODULE=on go
GOTESTCMD=GO111MODULE=on gotest
LOCALCMD=/usr/local

linters-install:
	@golangci-lint --version >/dev/null 2>&1 || { \
		echo "installing linting tools..."; \
		curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s v1.21.0; \
	}

lint: linters-install
	 $(LOCALCMD)/bin/golangci-lint run

fmt:
	$(GOCMD) fmt ./...

vet:
	$(GOCMD) vet ./.

test:
	$(GOTESTCMD) -v -cover -coverprofile=cover.out

cover:
	$(GOCMD) tool cover -html=cover.out

bench:
	$(GOTESTCMD) -bench=. -benchmem ./...
