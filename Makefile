default: lint deps test

lint:
	go get github.com/golang/lint/golint
	golint ./...

test:
	godep go test ./...

deps:
	go get ./...
	go get google.golang.org/grpc
	go get github.com/golang/protobuf/{proto,protoc-gen-go}
	protoc --go_out=plugins=grpc:knife/proto/. knife/knife.proto
	godep save ./...
	go get github.com/golang/lint/golint
