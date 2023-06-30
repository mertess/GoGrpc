generate:
	protoc --go-grpc_out=. --go_out=. protos\wordsSeen.proto