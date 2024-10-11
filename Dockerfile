FROM golang:latest AS builder

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY cmd cmd/
COPY internal internal/

WORKDIR /app/cmd
RUN go build -o /app/app app.go

FROM alpine:3.12 AS runner

RUN apk update && apk add ca-certificates libc6-compat

COPY templates /templates
COPY --from=builder /app/app /app

CMD /app
