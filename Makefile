

all: proto

proto:
	protoc -I proto/ proto/horeb.proto --go_out=plugins=grpc:proto

package-arm:
	holo-build -f --format=pacman ./build/package/holo.toml

package-x86_64:
	holo-build -f --format=pacman ./build/package/holo.toml

.PHONY: all proto package-arm package-x86_64

