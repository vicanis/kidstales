FROM golang:1.21 AS builder

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY cmd cmd/
COPY internal internal/

WORKDIR /app/cmd
RUN go build -o /app/app app.go

FROM ubuntu:latest AS runner

RUN apt-get update && apt-get install -y ca-certificates

COPY templates /templates
COPY --from=builder /app/app /app

CMD /app
