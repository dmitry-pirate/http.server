FROM golang:alpine as go

RUN apk update && apk add --no-cache git

WORKDIR /app
COPY . /app

RUN go build -v main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=go /app/main .
COPY --from=go /app/.env .

EXPOSE 8080

CMD /app/main