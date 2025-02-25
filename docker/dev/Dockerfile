FROM golang:1.23.6 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /app/cmd/main ./cmd/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /app
COPY --from=builder /app/cmd/main /app/cmd/main
COPY --from=builder /app/config ./config
COPY --from=builder /app/migrations ./migrations

ENV GOENV=dev
EXPOSE 8080

CMD ["/app/cmd/main"]