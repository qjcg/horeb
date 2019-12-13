VERSION := $(shell git describe --tags)
VERSION_IMPORTPATH := github.com/qjcg/horeb/pkg/horeb.Version


.PHONY: all
all: install

.PHONY: install
install: proto
	go install -ldflags '-s -w -X $(VERSION_IMPORTPATH)=$(VERSION)' ./...

.PHONY: build_images
build_images: proto
	docker build --build-arg VERSION=$(VERSION) --target horebd -t horebd .
	docker build --build-arg VERSION=$(VERSION) --target horebctl -t horebctl .

.PHONY: proto
proto: proto/horeb.proto
	protoc -I proto/ proto/horeb.proto --go_out=plugins=grpc:proto
