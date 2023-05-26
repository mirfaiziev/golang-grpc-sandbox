#build golang image
FROM golang:1.20-alpine as base

FROM base as dev

WORKDIR /server

COPY ./ ./

RUN apk update
RUN apk add --no-cache git
RUN apk add --no-cache make protobuf-dev

# todo: move it into requires
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

EXPOSE 8080

