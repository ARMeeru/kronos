# Kronos
Kronos is a simple command-line utility to test the response time of HTTP APIs. You can specify the API method, URL, headers, request body, and the acceptable response time threshold, and Kronos will make the HTTP request and report whether the response time was within the threshold.
## Usage
Build kronos using the following command

    go build -o kronos main.go

Ask for available flags

    ./kronos -help

Get the program running

`./kronos -csv api_info.csv` OR `./kronos -csv api_info.csv -threshold 3`

Sample response

    Response time for GET https://api.ipify.org: 1.978596s
    
    Test passed for GET https://api.ipify.org.
    
    Response time for POST https://reqbin.com/echo/post/json: 38.443ms
    
    Test passed for POST https://reqbin.com/echo/post/json.

## Contributing
If you find a bug or have a feature request, please open an issue on GitHub. Pull requests are also welcome.
## To improve
- Error handling

-  ~~User feedback~~

- ~~Testing~~