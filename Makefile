.PHONY: test
test:
	 go test -v -failfast ./...

.PHONY: coverage 
coverage:
	 go test -coverprofile=cover.out  ./...
	 go tool cover -html=cover.out -o cover.html
	 open cover.html
