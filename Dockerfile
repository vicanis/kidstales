FROM golang:latest AS build

COPY go.* /app/

COPY . /app/
WORKDIR /app/cmd
RUN go build -o /app/app app.go

FROM alpine:latest

RUN apk update && apk add ca-certificates libc6-compat

COPY templates /templates
COPY --from=build /app/app /app

CMD /app
