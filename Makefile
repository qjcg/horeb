version := $(shell git describe --tags)
version_importpath := github.com/qjcg/horeb/internal/horeb.Version
img := horeb
img_latest := $(img):latest
img_versioned := $(img):$(version)


.PHONY: all
all: install

.PHONY: install
install:
	go install -ldflags '-s -w -X $(version_importpath)=$(version)' ./...

.PHONY: docker-build
docker-build:
	docker build --build-arg VERSION=$(version) --target horeb -t $(img_latest) -t $(img_versioned) .

.PHONY: docker-run
docker-run:
	docker run --rm $(img_latest)

.PHONY: docker-run-interactive
docker-run-interactive:
	docker run --rm -it --entrypoint sh $(img_latest)
