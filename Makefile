.PHONY: fmt check-fmt lint vet test download install

GO_PKGS   := $(shell go list -f {{.Dir}} ./...)

fmt:
	@go list -f {{.Dir}} ./... | xargs -I{} gofmt -w -s {}

lint:
	@which golangci-lint > /dev/null || ( echo golangci-lint not found. run "make install" and try again.; false; )
	@golangci-lint run

test:
	@go test -race -v $(GO_FLAGS) -count=1 $(GO_PKGS)

download:
	@echo Download go.mod dependencies
	@go mod download

install: download
	@echo Installing tools from tools.go
	@cat tools.go | grep _ | awk -F'"' '{print $$2}' | xargs -tI % go install %
