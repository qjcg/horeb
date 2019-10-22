PACKAGE := github.com/qjcg/horeb/pkg/horeb
VERSION := $(shell git describe --tags)


all: install

install: proto
	go install -ldflags '-s -w -X $(PACKAGE).Version=$(VERSION)' ./...

proto: proto/horeb.proto
	protoc -I proto/ proto/horeb.proto --go_out=plugins=grpc:proto

.PHONY: all proto install
