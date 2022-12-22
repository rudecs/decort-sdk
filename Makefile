.PHONY: lint
lint:
	golangci-lint run --timeout 600s

.DEFAULT_GOAL := lint