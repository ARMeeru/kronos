package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

// The client field specifies the settings for the HTTP client, including the maximum time allowed for a request to complete.
type Kronos struct {
	client *http.Client
}

type APIInfo struct {
	Method string
	URL    string
	// The Headers field is a map that associates string keys with string values
	Headers     map[string]string
	RequestBody string
}

// By returning a pointer here, the function ensures that any changes made to the Kronos instance will be reflected in the original instance, rather than creating a copy of the instance.
func NewKronos() *Kronos {
	return &Kronos{
		client: &http.Client{Timeout: 5 * time.Second},
	}
}

func (k *Kronos) Run(apiInfo *APIInfo, threshold time.Duration) {
	req, err := http.NewRequest(apiInfo.Method, apiInfo.URL, nil)
	if err != nil {
		fmt.Printf("Error creating HTTP request for %s %s: %v\n", apiInfo.Method, apiInfo.URL, err)
		return
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
		fmt.Printf("Error sending HTTP request for %s %s: %v\n", apiInfo.Method, apiInfo.URL, err)
		return
	}

	defer resp.Body.Close()

	fmt.Printf("Response time for %s %s: %v\n", apiInfo.Method, apiInfo.URL, elapsed)

	if elapsed > threshold {
		fmt.Printf("Test failed: response time for %s %s exceeded threshold of %v.\n", apiInfo.Method, apiInfo.URL, threshold)
	} else {
		fmt.Printf("Test passed for %s %s.\n", apiInfo.Method, apiInfo.URL)
	}
}

func readAPIInfoFromCSV(filename string) ([]*APIInfo, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	r := csv.NewReader(file)
	var apiInfos []*APIInfo

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		apiInfo := &APIInfo{
			Method:      record[0],
			URL:         record[1],
			Headers:     parseKeyValueString(record[2]),
			RequestBody: record[3],
		}

		apiInfos = append(apiInfos, apiInfo)
	}

	return apiInfos, nil
}

func parseKeyValueString(s string) map[string]string {
	headers := make(map[string]string)
	headerPairs := strings.Split(s, ",")
	for _, pair := range headerPairs {
		pairParts := strings.Split(pair, "=")
		if len(pairParts) == 2 {
			headers[pairParts[0]] = pairParts[1]
		}
	}
	return headers
}

func main() {
	csvFilename := flag.String("csv", "", "CSV file containing API info")
	thresholdStr := flag.String("threshold", "1", "Acceptable response time threshold (in seconds)")
	flag.Parse()

	if *csvFilename == "" {
		fmt.Println("Error: CSV file path cannot be empty")
		os.Exit(1)
	}

	apiInfos, err := readAPIInfoFromCSV(*csvFilename)
	if err != nil {
		fmt.Println("Error reading API info from CSV:", err)
		os.Exit(1)
	}

	threshold, err := strconv.Atoi(*thresholdStr)
	if err != nil {
		fmt.Println("Invalid threshold value, using default value of 1 seconds")
		threshold = 3
	}

	k := NewKronos()
	for _, apiInfo := range apiInfos {
		k.Run(apiInfo, time.Duration(threshold)*time.Second)
	}
}
