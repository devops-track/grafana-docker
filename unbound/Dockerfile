FROM golang:alpine AS build
LABEL maintainer "Thomas Sänger <thomas@gecko.space>"

ENV CGO_ENABLED=0

RUN apk add --no-cache \
  ca-certificates \
  git

RUN go install -v github.com/letsencrypt/unbound_exporter@latest


FROM scratch
COPY --from=build /go/bin/* /

VOLUME /certs

CMD [ "/unbound_exporter", "-h" ]

EXPOSE 9167/tcp