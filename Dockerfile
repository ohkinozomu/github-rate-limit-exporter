FROM golang:1.17-buster as builder
WORKDIR /app
COPY go.* ./
RUN go mod download
COPY . ./
RUN go build -mod=readonly -v -o github-rate-limit-exporter

FROM debian:buster-slim
RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
    ca-certificates && \
    rm -rf /var/lib/apt/lists/*
COPY --from=builder /app/github-rate-limit-exporter /app/github-rate-limit-exporter
CMD ["/app/github-rate-limit-exporter"]