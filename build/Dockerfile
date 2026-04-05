FROM golang:1.26.1 AS builder

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /out/htredirect ./cmd/htredirect

FROM alpine:3.20

WORKDIR /app

COPY --from=builder /out/htredirect /app/htredirect
COPY --from=builder /src/web/templates /app/web/templates

ENTRYPOINT ["/app/htredirect"]
