FROM golang:alpine as go

RUN apk update && apk add --no-cache git ca-certificates

WORKDIR /app
COPY . /app

RUN go build -v main.go

EXPOSE 8080

CMD /app/main