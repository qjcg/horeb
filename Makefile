all: proto

proto:
	protoc -I proto/ proto/horeb.proto --go_out=plugins=grpc:proto

.PHONY: all proto

