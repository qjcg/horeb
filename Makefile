VERSION := $(shell git describe --tags)
VERSION_IMPORTPATH := github.com/qjcg/horeb/pkg/horeb.Version


.PHONY: all
all: install

.PHONY: install
install:
	go install -ldflags '-s -w -X $(VERSION_IMPORTPATH)=$(VERSION)' ./...

.PHONY: docker-build
docker-build:
	docker build --build-arg VERSION=$(VERSION) --target horeb -t horeb:latest -t horeb:$(VERSION) .

.PHONY: clean
clean:
	-rm proto/horeb.pb.go
