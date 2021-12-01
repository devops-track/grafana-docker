FROM golang:1.16-alpine
WORKDIR /app
COPY cloudflare.go cloudflare.go
COPY main.go main.go
COPY prometheus.go prometheus.go
COPY go.mod go.mod
COPY go.sum go.sum
RUN go get -d -v
RUN apk update && apk add ca-certificates
COPY --from=builder /app/cloudflare_exporter cloudflare_exporter
ENV CF_API_KEY=1f777f3e7f53d44493880354e6059a09d1893
ENV CF_API_EMAIL=ulisses.donato@cherokee.com.br
ENV CF_API_TOKEN=xTmdAwKuaESFvIU83ic7wZy_bmXP8MBulk-nPhjW

ENTRYPOINT [ "./cloudflare_exporter" ]