FROM golang:1.19-alpine AS builder
WORKDIR /go/src/app
COPY . .
ARG VERSION
ARG VERSION_IMPORTPATH=main.Version
RUN CGO_ENABLED=0 go install -ldflags "-s -w -X $VERSION_IMPORTPATH=$VERSION" ./...

##########

# Use alpine as base to allow interactive use from a shell session.
FROM alpine:latest AS horeb
USER 1001
COPY --from=builder /go/bin/horeb /usr/bin/
ENTRYPOINT ["/usr/bin/horeb"]
