GO := go1.12beta2
MODULE := $(shell go list -m)
VERSION := $(shell git describe --tags)
PKGVER := $(shell git describe --tags --abbrev=0 | tr -d v)

OUTDIR := $(PWD)/build/package
BINARIES := horeb horebd horebctl
GOARCHES := amd64 arm
ARCHARCHES := x86_64 armv7h


all: package

proto:
	protoc -I proto/ proto/horeb.proto --go_out=plugins=grpc:proto

build: build-arm build-amd64

build-arm:
	@echo Compiling: arm
	@$(foreach bin,$(BINARIES),env GOARCH=arm GOARM=7 $(GO) build -ldflags '-X $(MODULE).Version=$(VERSION) -s -w' -o $(OUTDIR)/$(bin)-arm_$(VERSION) ./cmd/$(bin); )

build-amd64:
	@echo Compiling: amd64
	@$(foreach bin,$(BINARIES),$(GO) build -ldflags '-X $(MODULE).Version=$(VERSION) -s -w' -o $(OUTDIR)/$(bin)-amd64_$(VERSION) ./cmd/$(bin); )

compress: build
	@echo Compressing binaries
	@upx $(wildcard $(OUTDIR)/horeb*)

# The sigil tool renders Go template files.
# See https://github.com/gliderlabs/sigil
package-templates:
	@echo Generating holo templates
	@$(foreach arch,$(ARCHARCHES),sigil -f templates/holo.toml.tmpl Architecture=$(arch) Version=$(PKGVER) Binaries='' > build/package/holo-$(arch).toml; )

package: compress package-templates
	@echo Building holo packages
	@$(foreach arch,$(ARCHARCHES),holo-build -f --format=pacman ./build/package/holo-$(arch).toml; )

clean:
	rm -f $(wildcard $(OUTDIR)/horeb*)

.PHONY: all proto build build-arm build-amd64 compress package-templates package clean
