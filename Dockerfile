# Build executable binary
FROM golang:1.22.4-alpine3.20 AS builder
LABEL maintainer="Alfian Akmal Hanantio<amalhanaja@gmail.com>"
RUN apk update && apk add --no-cache git && apk add --no-cach bash && apk add build-base
WORKDIR /app
COPY ./go.mod ./
COPY ./go.sum ./
RUN go mod download
COPY . .
RUN go build -o /cli cmd/cli/main.go

# build small iamge
FROM alpine:3.20.0
WORKDIR /app
COPY --from=builder cli .
CMD [ "./cli" ]