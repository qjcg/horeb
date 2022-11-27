# Use alpine as base to allow interactive use from a shell session.
FROM alpine:latest
USER 1001
ENTRYPOINT ["/usr/bin/horeb"]
COPY horeb /usr/bin/
