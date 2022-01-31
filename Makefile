.PHONY: test
test:
	go clean -testcache
	go test ./...

.PHONY: test-of-tests
test-of-tests:
	./test_of_tests.sh
