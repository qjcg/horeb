GO := go1.12beta2
MODULE := $(shell go list -m)
VERSION := $(shell git describe --tags)
PKGVER := $(shell git describe --tags --abbrev=0 | tr -d v)	# Used in holo packages

OUTDIR := $(PWD)/build/package
OUTDIR_ARM := $(OUTDIR)/arm
OUTDIR_AMD64 := $(OUTDIR)/amd64

BINARIES := $(shell ls cmd)
GOARCHES := amd64 arm
ARCHARCHES := x86_64 armv7h					# Used in holo packages


all: package

proto:
	protoc -I proto/ proto/horeb.proto --go_out=plugins=grpc:proto

build: build-arm build-amd64

build-arm:
	@echo Compiling: arm
	@mkdir -p $(OUTDIR_ARM)
	@$(foreach bin,$(BINARIES),env GOARCH=arm GOARM=7 $(GO) build -ldflags '-X $(MODULE).Version=$(VERSION) -s -w' -o $(OUTDIR_ARM)/$(bin) ./cmd/$(bin); )

build-amd64:
	@echo Compiling: amd64
	@mkdir -p $(OUTDIR_AMD64)
	@$(foreach bin,$(BINARIES),$(GO) build -ldflags '-X $(MODULE).Version=$(VERSION) -s -w' -o $(OUTDIR_AMD64)/$(bin) ./cmd/$(bin); )

compress: build
	@echo Compressing binaries
	@upx $(wildcard $(OUTDIR)/*/horeb*)

# The sigil tool renders Go template files.
# See https://github.com/gliderlabs/sigil
package-templates:
	@echo Generating holo templates
	@sigil -f templates/holo.toml.tmpl Architecture="x86_64" Version=$(PKGVER) Binaries='horeb,horebd,horebctl' > $(OUTDIR_AMD64)/holo.toml
	@sigil -f templates/holo.toml.tmpl Architecture="armv7h" Version=$(PKGVER) Binaries='horeb,horebd,horebctl' > $(OUTDIR_ARM)/holo.toml

package: compress package-templates
	@echo Building holo packages

	@# Link files referenced in holo.toml file.
	@$(foreach d,$(OUTDIR_AMD64) $(OUTDIR_ARM),ln -s $(PWD)/LICENSE $(PWD)/init/* $(d); )

	@cd $(OUTDIR_AMD64); holo-build -f --format=pacman holo.toml
	@cd $(OUTDIR_ARM); holo-build -f --format=pacman holo.toml

clean:
	@echo Removing build artifacts
	@rm -rf $(OUTDIR)

.PHONY: all proto build build-arm build-amd64 compress package-templates package clean
