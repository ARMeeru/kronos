# Kronos
Kronos is a simple command-line utility to test the response time of HTTP APIs. You can specify the API method, URL, headers, request body, and the acceptable response time threshold, and Kronos will make the HTTP request and report whether the response time was within the threshold.
## Installation
To use Kronos, you need to have Go installed on your system. You can download and install the latest version of Go from the official [website](https://golang.org/dl).

Once you have installed Go, you can install Kronos by running the following command:

    go get github.com/ARMeeru/kronos

This will download and install Kronos and its dependencies.
## Usage
To use Kronos, simply run the `kronos` command in your terminal, and follow the prompts to enter the API information:

    $ ./kronos OR go run main.go
    Enter API method (GET/POST/PUT/DELETE): GET
    Enter API headers (in key=value format, separated by commas): Authorization=Bearer <token>, Content-Type=application/json
    Enter API URL: https://api.example.com/v1/users
    Enter API request body (in key=value format, separated by commas): email=john@example.com
    Enter acceptable response time threshold (in seconds): 1
    Response time: 534.259ms
    Test passed.

## Contributing
If you find a bug or have a feature request, please open an issue on GitHub. Pull requests are also welcome.
## To improve
- Error handling
- User feedback
