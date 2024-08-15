FROM golang:1.22.5 AS builder

WORKDIR /app

COPY go.mod go.sum Makefile ./

RUN make deps

COPY . .

RUN make build-binary

FROM alpine:3.20.2 AS runner

RUN apk add --no-cache libc6-compat

WORKDIR /app

COPY --from=builder /app/build/output/main /app
COPY --from=builder /app/config config
COPY --from=builder /app/migrations migrations

CMD ["./main"]