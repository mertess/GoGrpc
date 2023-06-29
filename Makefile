generate:
	mkdir .\gen;
	protoc --go-grpc_out=gen protos\wordsSeen.proto