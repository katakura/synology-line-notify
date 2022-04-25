# syntax=docker/dockerfile:1

FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod ./
#COPY go.sum ./
RUN go mod download

COPY notify.go ./

RUN go build -o /notify

EXPOSE 8080

CMD [ "/notify" ]
