version := $(shell git describe --tags)
version_importpath := main.Version
ldflags := -ldflags '-s -w -X $(version_importpath)=$(version)'
img := horeb
img_latest := $(img):latest
img_versioned := $(img):$(version)


.PHONY: all
all: install

.PHONY: build
build:
	goreleaser build --rm-dist

.PHONY: build-snapshot
build-snapshot:
	goreleaser build --rm-dist --snapshot

.PHONY: release
release:
	goreleaser release --rm-dist

.PHONY: release-snapshot
release-snapshot:
	goreleaser release --rm-dist --snapshot

.PHONY: install
install:
	GOAMD64=v3 go install $(ldflags) ./cmd/horeb
	upx $(GOBIN)/horeb

.PHONY: clean
clean:
	rm -rf horeb coverprofile.out dist

.PHONY: test
test: test-integration

.PHONY: test-unit
test-unit:
	go test -cover ./...

.PHONY: test-integration
test-integration:
	go test -tags integration -v ./...

.PHONY: cover
cover:
	go test -tags integration -coverprofile coverprofile.out ./...
	go tool cover -func coverprofile.out
	go tool cover -html coverprofile.out

.PHONY: tag
tag:
	git tag -am "$(t)" $(t)
	git push --tags


# Docker

.PHONY: docker-build
docker-build:
	docker build --build-arg VERSION=$(version) --target horeb -t $(img_latest) -t $(img_versioned) .

.PHONY: docker-run
docker-run:
	docker run --rm $(img_latest)

.PHONY: docker-run-interactive
docker-run-interactive:
	docker run --rm -it --entrypoint sh $(img_latest)
