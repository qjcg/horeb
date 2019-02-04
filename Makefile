all: proto

proto:
	protoc -I proto/ proto/horeb.proto --go_out=plugins=grpc:proto

package:
	holo-build -f --format=pacman ./build/package/holo.toml

.PHONY: all proto

