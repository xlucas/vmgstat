NAME = vmgstat
DEPS = $(shell go list -f '{{range .Imports}}{{.}} {{end}}' ./... | tr ' ' '\n' | grep "github.com" | grep -v $(NAME) | tr '\n' ' ')

deps:
	go get -d -v $(DEPS)

test: deps
	go test -v ./...

.PHONY: deps test
