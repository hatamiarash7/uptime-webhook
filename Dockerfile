FROM alpine:3.18.0 as certs

RUN apk --update add ca-certificates

FROM debian:stretch-20220622

ARG DATE_CREATED
ARG APP_VERSION

LABEL org.opencontainers.image.created=$DATE_CREATED
LABEL org.opencontainers.version="$APP_VERSION"
LABEL org.opencontainers.image.title="uptime-webhook"
LABEL org.opencontainers.image.description="It's webhook handler for uptime.com"

COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

ADD bin/webhook /usr/sbin/webhook

CMD ["/usr/sbin/webhook"]
