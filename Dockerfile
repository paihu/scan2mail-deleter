FROM golang:1.12-alpine as builder

ADD . /work/
WORKDIR /work/
RUN apk add --no-cache git
RUN go build



FROM alpine:latest
COPY --from=builder /work/slack-delete-file-bot /bin/
RUN apk add --no-cache ca-certificates && update-ca-certificates
ENV SSL_CERT_FILE=/etc/ssl/certs/ca-certificates.crt
ENV SSL_CERT_DIR=/etc/ssl/certs/

ENTRYPOINT ["/bin/slack-delete-file-bot"]
