FROM golang:1.18.4-alpine3.15

RUN go install github.com/beego/bee/v2@latest

EXPOSE 8080