test:
	go test -timeout 2s ./...

proto:
	protoc --go_out=./src/ \
    --go-grpc_out=./src/ \
    .proto/pegasus.proto

hot:
	air
