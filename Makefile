default: deps lint protobuf test

lint:
	go get github.com/golang/lint/golint
	golint ./...

test:
	godep go test ./...

protobuf:
	go get google.golang.org/grpc
	go get github.com/golang/protobuf/proto
	go get github.com/golang/protobuf/protoc-gen-go
	protoc --go_out=plugins=grpc:knife/proto/. knife/knife.proto

deps:
	godep save ./...
