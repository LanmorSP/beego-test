FROM golang:1.18.4-alpine3.15

WORKDIR /app

COPY . /app
RUN go install github.com/beego/bee/v2@latest
RUN go mod tidy

EXPOSE 8080