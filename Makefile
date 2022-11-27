version := $(shell git describe --tags)
version_importpath := main.Version
ldflags := -ldflags '-s -w -X $(version_importpath)=$(version)'
img := horeb
img_latest := $(img):latest
img_versioned := $(img):$(version)


.PHONY: all
all: install

.PHONY: build-snapshot
build-snapshot:
	goreleaser build --rm-dist --snapshot

.PHONY: release-snapshot
release-snapshot:
	goreleaser release --rm-dist --snapshot

.PHONY: install
install:
	go install $(ldflags) ./cmd/horeb
	upx $(GOBIN)/horeb

.PHONY: clean
clean:
	rm -rf horeb coverprofile.out dist

.PHONY: test
test:
	go test -cover ./...

.PHONY: testall
testall:
	go test -cover -tags integration -v ./...

.PHONY: cover
cover:
	go test -tags integration -coverprofile coverprofile.out ./...
	go tool cover -func coverprofile.out
	go tool cover -html coverprofile.out

.PHONY: tag
tag: tag_next := $(shell svu next)
tag:
	git tag -am "$(tag_next)" $(tag_next)
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
