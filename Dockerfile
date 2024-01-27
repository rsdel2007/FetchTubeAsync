FROM golang:alpine3.18

WORKDIR /app
ENV CGO_ENABLED=1

RUN apk add --no-cache \
    # Important: required for go-sqlite3
    gcc \
    # Required for Alpine
    musl-dev
COPY . .
RUN go mod download


EXPOSE 8081

CMD ["go", "run", "cmd/main.go"]