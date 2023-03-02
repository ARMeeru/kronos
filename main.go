package main

import (
	"bytes"
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
	"time"
)

func main() {
	// Parse command-line arguments
	csvFilename := flag.String("csv", "api_info.csv", "CSV file containing API info")
	thresholdStr := flag.String("threshold", "1", "Acceptable response time threshold (in seconds)")
	flag.Parse()

	// Parse the threshold value
	threshold, err := time.ParseDuration(*thresholdStr + "s")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid threshold value: %v\n", err)
		threshold, _ = time.ParseDuration("1s") // Default threshold of 1 second
	}

	// Open the CSV file
	csvFile, err := os.Open(*csvFilename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening CSV file: %v\n", err)
		os.Exit(1)
	}
	defer csvFile.Close()

	// Parse the CSV data
	apiData, err := parseCsvData(csvFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing CSV data: %v\n", err)
		os.Exit(1)
	}

	// Set up a worker pool to process API requests concurrently
	numWorkers := 5
	var wg sync.WaitGroup
	wg.Add(numWorkers)
	apiDataChan := make(chan *apiInfo)
	for i := 0; i < numWorkers; i++ {
		go func() {
			defer wg.Done()

			// Loop through the API data channel and make requests
			for apiData := range apiDataChan {
				// Make the API request
				start := time.Now()
				resp, err := makeRequest(apiData)
				if err != nil {
					fmt.Fprintf(os.Stderr, "Error making request: %v\n", err)
					continue
				}
				responseTime := time.Since(start)

				// Check the response time
				if responseTime > threshold {
					fmt.Fprintf(os.Stderr, "Request to %s took too long (%v)\n", apiData.Url, responseTime)
					os.Exit(1)
				} else {
					fmt.Printf("Request to %s completed in %v\n", apiData.Url, responseTime)
				}

				// Close the response body
				_, _ = io.Copy(ioutil.Discard, resp.Body)
				resp.Body.Close()
			}
		}()
	}

	// Push the API data to the channel
	for _, data := range apiData {
		apiDataChan <- data
	}
	close(apiDataChan)

	// Wait for all workers to finish
	wg.Wait()

}

// Structure to hold API information
type apiInfo struct {
	Method string
	Url    string
	Body   string
	Header http.Header
}

// Parse CSV data into an array of API info structs
func parseCsvData(csvFile io.Reader) ([]*apiInfo, error) {
	var apiData []*apiInfo

	// Parse the CSV data
	csvReader := csv.NewReader(csvFile)
	csvData, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	// Get the headers from the CSV data
	headers := csvData[0]

	// Loop through the rows in the CSV data
	for i, row := range csvData[1:] {
		// Check that the row has the correct number of columns
		if len(row) != len(headers) {
			return nil, errors.New(fmt.Sprintf("Incorrect number of columns in row %d", i+1))
		}

		// Create a new API info struct
		info := &apiInfo{
			Method: row[0],
			Url:    row[1],
			Body:   row[2],
			Header: make(http.Header),
		}

		// Parse the headers
		for j := 3; j < len(headers); j++ {
			headerName := headers[j]
			headerValue := row[j]
			if headerValue != "" {
				info.Header.Add(headerName, headerValue)
			}
		}

		// Add the API info to the array
		apiData = append(apiData, info)
	}

	return apiData, nil
}

// Make an HTTP request and return the response and an error, if any
func makeRequest(apiData *apiInfo) (*http.Response, error) {
	// Create the request
	req, err := http.NewRequest(apiData.Method, apiData.Url, bytes.NewBufferString(apiData.Body))
	if err != nil {
		return nil, err
	}

	// Add the headers to the request
	req.Header = apiData.Header

	// Make the request
	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
