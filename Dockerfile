# Stage 1: build
FROM golang:1.23-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o /bin/api ./cmd/api

# Stage 2: run
FROM alpine:3.21

RUN apk add --no-cache ca-certificates tzdata

COPY --from=builder /bin/api /bin/api
COPY --from=builder /app/docs /docs
COPY --from=builder /app/migrations /migrations

EXPOSE 8080

ENTRYPOINT ["/bin/api"]
