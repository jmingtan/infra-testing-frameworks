# Serverspec example

Sample test cases for evaluating [Serverspec](https://serverspec.org) using Docker.

To see an example of a failing test case, run:

    docker-compose up serverspec-failure
    
To get a successful test, run:

    docker-compose up serverspec-success

Afterwards, you can clean up the environment with:

    docker-compose down
