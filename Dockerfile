FROM golang:1.20-alpine AS builder

WORKDIR /app
COPY . .
RUN go build -o gochat .

FROM scratch
ENV PORT=1337

COPY --from=builder /app/gochat /gochat

ENTRYPOINT ["/gochat"]
