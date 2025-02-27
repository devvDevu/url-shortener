FROM golang:1.23.6 AS builder

WORKDIR /url-shortener
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /url-shortener/cmd/main ./cmd/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /url-shortener
COPY --from=builder /url-shortener/cmd/main /url-shortener/cmd/main
COPY --from=builder /url-shortener/config ./config
COPY --from=builder /url-shortener/migrations ./migrations
COPY --from=builder /url-shortener/docs ./docs

ENV GOENV=dev
EXPOSE 8080

CMD ["/url-shortener/cmd/main"]