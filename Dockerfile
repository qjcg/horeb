FROM golang:1.13-alpine AS builder
WORKDIR /go/src/app
COPY [".", "."]
ARG PACKAGE=github.com/qjcg/horeb/pkg/horeb
RUN apk add --no-cache upx
# TODO: Set version via -ldflags!
RUN \
	CGO_ENABLED=0 go install -ldflags '-s -w' ./... && \
	upx /go/bin/*

##########

FROM scratch AS horebd
USER 1337
COPY --from=builder /go/bin/horebd /usr/bin/horebd
EXPOSE 9999
ENTRYPOINT ["/usr/bin/horebd"]

##########

# Use alpine as base to allow interactive use from a shell session.
FROM alpine:latest AS horebctl
USER 1337
COPY --from=builder /go/bin/horebctl /usr/bin/horebctl
COPY --from=builder /go/bin/horeb /usr/bin/horeb
ENTRYPOINT ["/bin/sh"]
