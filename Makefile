GO := go1.12beta2
MODULE := github.com/qjcg/horeb
VERSION := $(shell git describe --tags)

BINARIES := horeb horebd horebctl
GOARCHES := amd64 arm
ARCHARCHES := x86_64 armv7h


all: proto

proto:
	protoc -I proto/ proto/horeb.proto --go_out=plugins=grpc:proto

build: build-arm build-amd64

build-arm:
	$(foreach bin,$(BINARIES),env GOARCH=arm GOARM=7 $(GO) build -ldflags '-X $(MODULE).Version=$(VERSION) -s -w' -o $(bin)-arm_$(VERSION) ./cmd/$(bin); )

build-amd64:
	$(foreach bin,$(BINARIES),$(GO) build -ldflags '-X $(MODULE).Version=$(VERSION) -s -w' -o $(bin)-amd64_$(VERSION) ./cmd/$(bin); )

compress: build
	upx $(wildcard horeb*)

# The sigil tool renders Go template files.
# See: https://github.com/gliderlabs/sigil
package-templates:
	$(foreach arch,$(ARCHARCHES),sigil -f templates/holo.toml.tmpl Architecture=$(arch) Version=$(VERSION) Binaries='' > build/package/holo-arch.toml; )

package: compress
	$(foreach arch,,holo-build -f --format=pacman ./build/package/holo.toml; )

.PHONY: all proto build build-arm build-amd64 compress package-templates package
