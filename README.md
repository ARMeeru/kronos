# 1. Kronos
Kronos is a simple command-line utility to test the response time of HTTP APIs. You can specify the API method, URL, headers, request body, and the acceptable response time threshold, and Kronos will make the HTTP request and report whether the response time was within the threshold.
## 1.1. Usage
Build kronos using the following command

    go build -o kronos main.go

Ask for available flags

    ./kronos -help

Get the program running
`./kronos -csv api_info.csv` OR `./kronos -csv api_info.csv -threshold 3`

Sample response

    Request to https://reqbin.com/echo/post/json completed in 152.280167ms
    Request to https://api.ipify.org took too long (1.251576584s)

## 1.2. Contributing
If you find a bug or have a feature request, please open an issue on GitHub. Pull requests are also welcome.
## 1.3. To improve
- Implement concurrency
- ~~Error handling~~
- ~~User feedback~~
- Testing (files will be added properly later)