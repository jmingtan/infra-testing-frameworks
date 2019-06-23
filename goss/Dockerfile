FROM alpine:3.10.0
RUN apk add curl nginx
RUN mkdir -p /run/nginx
RUN curl -L https://github.com/aelsabbahy/goss/releases/download/v0.3.7/goss-linux-amd64 -o /usr/local/bin/goss
RUN chmod +x /usr/local/bin/goss
WORKDIR /usr/src/app
COPY docker-entrypoint.sh .
CMD ["./docker-entrypoint.sh"]