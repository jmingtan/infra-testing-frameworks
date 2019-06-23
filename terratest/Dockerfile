FROM golang:1.12.6-alpine3.10
RUN apk add nginx git gcc libc-dev
RUN mkdir -p /run/nginx
WORKDIR /usr/src/app
COPY . .
RUN go mod download
CMD ["go", "test", "-v"]