# goss example

Sample test cases for evaluating [goss](https://github.com/aelsabbahy/goss) using Docker.

To see goss in action, run:

    docker-compose up

We make use of the goss autoadd feature to validate the presence of
certain packages, and to verify that port 80 is listening. (see
`docker-entrypoint.sh` for details)

Afterwards, you can clean up the environment with:

    docker-compose down
