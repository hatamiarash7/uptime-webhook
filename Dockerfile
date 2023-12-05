FROM --platform=${BUILDPLATFORM:-linux/amd64} alpine:3.18.5 as certs

RUN apk --update add ca-certificates

FROM --platform=${BUILDPLATFORM:-linux/amd64} golang:1.21.5 as builder

ARG TARGETPLATFORM
ARG BUILDPLATFORM
ARG TARGETOS
ARG TARGETARCH

WORKDIR /app/
ADD . .

RUN GO111MODULE=on CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -ldflags="-s -w" -o webhook cmd/*.go

FROM --platform=${TARGETPLATFORM:-linux/amd64} scratch

ARG DATE_CREATED
ARG APP_VERSION
ENV APP_VERSION=$APP_VERSION

LABEL org.opencontainers.image.created=$DATE_CREATED
LABEL org.opencontainers.version="$APP_VERSION"
LABEL org.opencontainers.image.authors="Arash Hatami <info@arash-hatami.ir>"
LABEL org.opencontainers.image.vendor="hatamiarash7"
LABEL org.opencontainers.image.title="uptime-webhook"
LABEL org.opencontainers.image.description="It's a webhook handler for uptime.com"
LABEL org.opencontainers.image.source="https://github.com/hatamiarash7/uptime-webhook"
LABEL org.opencontainers.image.url="https://github.com/hatamiarash7/uptime-webhook"
LABEL org.opencontainers.image.documentation="https://github.com/hatamiarash7/uptime-webhook"

WORKDIR /app/

COPY --from=builder /app/webhook /app/webhook
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

EXPOSE 8080

ENTRYPOINT ["/app/webhook"]