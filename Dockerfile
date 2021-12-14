# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /commonsequence

COPY go.mod ./
COPY *.go ./
COPY /test ./

RUN go build -o commonsequence
