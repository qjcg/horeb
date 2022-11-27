version := $(shell git describe --tags)
version_importpath := main.Version
img := horeb
img_latest := $(img):latest
img_versioned := $(img):$(version)


.PHONY: all
all: install

.PHONY: build
build:
	go build -ldflags '-s -w -X $(version_importpath)=$(version)' ./cmd/horeb
	upx horeb

.PHONY: install
install:
	go install -ldflags '-s -w -X $(version_importpath)=$(version)' ./cmd/horeb
	upx $(GOBIN)/horeb

.PHONY: clean
clean:
	rm -rf horeb

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
