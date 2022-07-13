FROM alpine
WORKDIR /app
ADD web web
ADD build build
ENTRYPOINT build/rn-agent-go
