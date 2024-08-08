.PHONY: tests
tests:
	go test --race ./...

.PHONE: test-by-name
test-by-name:
	go test --race -run ${name} ./...
