FROM golang:1.13-alpine AS builder
WORKDIR /go/src/app
COPY [".", "."]
ARG VERSION
ARG VERSION_IMPORTPATH=github.com/qjcg/horeb/pkg/horeb.Version
RUN apk add --no-cache upx
RUN \
	CGO_ENABLED=0 go install -ldflags "-s -w -X $VERSION_IMPORTPATH=$VERSION" ./... && \
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
ENV HOREBCTL_HOST=horebd
# Sleep for a day. Used to keep container alive, and ready to be attached to
# with docker-compose.
ENTRYPOINT ["/bin/sleep", "1d"]
