VERSION := $(shell git describe --tags)
VERSION_IMPORTPATH := github.com/qjcg/horeb/pkg/horeb.Version


.PHONY: all
all: install

.PHONY: install
install: proto
	go install -ldflags '-s -w -X $(VERSION_IMPORTPATH)=$(VERSION)' ./...

.PHONY: build_images
build_images: proto
	docker build --build-arg VERSION=$(VERSION) --target horebd -t horebd:latest -t horebd:$(VERSION) .
	docker build --build-arg VERSION=$(VERSION) --target horebctl -t horebctl:latest -t horebctl:$(VERSION) .

# proto is simply an alias.
.PHONY: proto
proto: proto/horeb.pb.go

proto/horeb.pb.go: proto/horeb.proto
	protoc -I proto/ proto/horeb.proto --go_opt=paths=source_relative --go_out=plugins=grpc:proto

.PHONY: clean
clean:
	-rm proto/horeb.pb.go
