FROM golang:1.18-alpine
RUN apk add build-base

RUN go install github.com/go-delve/delve/cmd/dlv@latest

ENV APISERVER_CERT_LOCATION /app/crt/vkubelet-mock-0-crt.pem
ENV APISERVER_KEY_LOCATION /app/crt/vkubelet-mock-0-key.pem

WORKDIR /app
COPY . .
ENTRYPOINT build/rn-agent-go
