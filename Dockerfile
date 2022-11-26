FROM golang:1.19-alpine AS builder
WORKDIR /go/src/app
COPY . .
ARG VERSION
ARG VERSION_IMPORTPATH=github.com/qjcg/horeb/pkg/horeb.Version
RUN CGO_ENABLED=0 go install -ldflags "-s -w -X $VERSION_IMPORTPATH=$VERSION" ./...

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
# Loop and sleep. Used to keep container alive, and ready to be attached to
# with docker-compose.
CMD ["sh", "-c", "'while horebctl; do sleep 1s; done'"]
