FROM golang:1.19.1-alpine

RUN apk update && apk add git

RUN mkdir /app

WORKDIR /app

ADD . /app

RUN go install golang.org/x/tools/gopls@latest
