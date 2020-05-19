# syntax = docker/dockerfile:experimental

FROM golang:1.13

WORKDIR /app

COPY ./ ./

RUN make tools

RUN make fmt

RUN make build