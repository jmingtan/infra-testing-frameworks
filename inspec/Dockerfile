FROM ubuntu:18.04
RUN apt-get update && apt-get install -y curl nginx && rm -rf /var/lib/apt/lists/*
RUN mkdir -p /run/nginx
RUN curl -o /tmp/inspec.deb https://packages.chef.io/files/stable/inspec/3.9.3/ubuntu/18.04/inspec_3.9.3-1_amd64.deb && \
  dpkg -i /tmp/inspec.deb && \
  rm /tmp/inspec.deb
WORKDIR /usr/src/app
COPY . .
CMD ["./docker-entrypoint.sh"]
