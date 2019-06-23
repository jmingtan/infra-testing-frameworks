# terratest example

Sample test cases for evaluating
[terratest](https://github.com/gruntwork-io/terratest) using
Docker. Note that this test only uses a small portion of terratest's
API and does not represent it's overall capabilities.

To see terratest in action, run:

    docker-compose up
    
To see an example of a test failure, kill the web container:

    docker-compose kill web
    docker-compose up --no-deps terratest

Afterwards, you can clean up the environment with:

    docker-compose down
