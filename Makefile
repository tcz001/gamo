default: deps lint test

lint:
	golint ./...

test:
	godep go test ./...

deps:
	godep save ./...
	go get github.com/golang/lint/golint
