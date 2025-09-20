FROM golang:latest AS builder

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY cmd cmd/
COPY internal internal/

WORKDIR /app/cmd
RUN go build -o /app/app app.go

FROM ubuntu:22.04 AS runner

RUN apt-get update && \
    apt-get install -y gnupg2 && \
    apt-key adv --keyserver keyserver.ubuntu.com --recv-keys 871920D1991BC93C && \
    apt-get update && \
    apt-get install -y ca-certificates

COPY templates /templates
COPY --from=builder /app/app /app

CMD ["/app"]

