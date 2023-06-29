generate:
	mkdir .\gen;
	protoc --go_out=paths=source_relative:.\gen .\protos\wordsSeen.proto