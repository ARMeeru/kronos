package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type apiInfo struct {
	Method      string
	Headers     map[string]string
	RequestBody string
}

type Kronos struct {
	reader *bufio.Reader
	client *http.Client
}

func NewKronos() *Kronos {
	return &Kronos{
		reader: bufio.NewReader(os.Stdin),
		client: &http.Client{Timeout: 5 * time.Second},
	}
}

func (k *Kronos) Run() {
	apiInfo, err := k.getAPIInfoFromUser()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error getting API information:", err)
		os.Exit(1)
	}

	url, err := k.getURLFromUser()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error getting API URL:", err)
		os.Exit(1)
	}

	threshold, err := k.getThresholdFromUser()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error getting threshold:", err)
		os.Exit(1)
	}

	req, err := http.NewRequest(apiInfo.Method, url, nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating HTTP request:", err)
		os.Exit(1)
	}

	for key, value := range apiInfo.Headers {
		req.Header.Set(key, value)
	}

	if apiInfo.RequestBody != "" {
		req.Body = ioutil.NopCloser(strings.NewReader(apiInfo.RequestBody))
	}

	start := time.Now()
	resp, err := k.client.Do(req)
	elapsed := time.Since(start)

	if err != nil {
		fmt.Fprintln(os.Stderr, "Error sending HTTP request:", err)
		os.Exit(1)
	}

	defer resp.Body.Close()

	fmt.Printf("Response time: %v\n", elapsed)

	if elapsed > threshold {
		fmt.Printf("Test failed: response time exceeded threshold of %v.\n", threshold)
	} else {
		fmt.Println("Test passed.")
	}
}

func (k *Kronos) getAPIInfoFromUser() (apiInfo, error) {
	var info apiInfo

	fmt.Print("Enter API method (GET/POST/PUT/DELETE): ")
	method, err := k.reader.ReadString('\n')
	if err != nil {
		return info, err
	}
	method = strings.TrimSpace(method)
	if method == "" {
		return info, fmt.Errorf("method cannot be empty")
	}

	fmt.Print("Enter API headers (in key=value format, separated by commas): ")
	headersStr, err := k.reader.ReadString('\n')
	if err != nil {
		return info, err
	}
	headersStr = strings.TrimSpace(headersStr)
	headers := make(map[string]string)
	if headersStr != "" {
		headersArr := strings.Split(headersStr, ",")
		for _, header := range headersArr {
			headerKV := strings.Split(header, "=")
			if len(headerKV) == 2 {
				headers[headerKV[0]] = headerKV[1]
			}
		}
	}

	fmt.Print("Enter API request body (in key=value format, separated by commas): ")
	requestBody, err := k.reader.ReadString('\n')
	if err != nil {
		return info, err
	}
	requestBody = strings.TrimSpace(requestBody)

	info.Method = method
	info.Headers = headers
	info.RequestBody = requestBody

	return info, nil
}

func (k *Kronos) getURLFromUser() (string, error) {
	fmt.Print("Enter API URL: ")
	url, err := k.reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	url = strings.TrimSpace(url)
	if url == "" {
		return "", fmt.Errorf("URL cannot be empty")
	}
	return url, nil
}

func (k *Kronos) getThresholdFromUser() (time.Duration, error) {
	fmt.Print("Enter acceptable response time threshold (in seconds): ")
	thresholdStr, err := k.reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	thresholdStr = strings.TrimSpace(thresholdStr)
	if thresholdStr == "" {
		return 0, fmt.Errorf("threshold cannot be empty")
	}

	threshold, err := strconv.Atoi(thresholdStr)
	if err != nil {
		fmt.Println("Invalid input, using default threshold of 3 seconds.")
		return 3 * time.Second, nil
	}

	return time.Duration(threshold) * time.Second, nil
}

func main() {
	k := NewKronos()
	k.Run()
}
