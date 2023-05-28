FROM alpine:3.18.0 as certs

RUN apk --update add ca-certificates

FROM debian:stretch-20220622

COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

ADD bin/webhook /usr/sbin/webhook

CMD ["/usr/sbin/webhook"]
