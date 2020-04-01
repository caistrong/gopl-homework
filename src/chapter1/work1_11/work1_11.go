package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

// run result: caught errors and goroutine blocked

// why caught error `read: connection reset by peer`
// see related issuse: https://github.com/golang/go/issues/20960

// why blocked?
// we have 500 website, so we loop 500 times for `<-ch`
// but some of that 500 website didn't get response
// it cause that we push data to `ch` less 500 times

// run command:
// go run ./src/chapter1/work1_11/work1_11.go
func main() {
	startTime := time.Now()
	ch := make(chan string)
	websiteList := readCsvFile("./src/chapter1/work1_11/top500Domains.csv")
	for _, url := range websiteList {
		go fetch(url, ch)
	}

	for range websiteList {
		fmt.Println(<-ch)
	}
	costTime := time.Since(startTime).Seconds()
	fmt.Printf("%.2fs elasped\n", costTime)
}

func fetch(url string, ch chan<- string) {
	startTime := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	nBytes, copyErr := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if copyErr != nil {
		ch <- fmt.Sprintf("while reading url:%s, catch err: %v", url, copyErr)
		return
	}
	costTime := time.Since(startTime).Seconds()
	ch <- fmt.Sprintf("%.2fs\t%7d\t%s", costTime, nBytes, url)
}

func readCsvFile(filePath string) []string {
	// Load a csv file.
	websiteList := []string{}

	f, err := os.Open(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "open err:%v", err)
		return websiteList
	}

	// Create a new reader.
	r := csv.NewReader(f)
	for {
		record, err := r.Read()
		// Stop at EOF.
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Fprintf(os.Stderr, "read err:%v", err)
			break

		}
		websiteList = append(websiteList, "http://"+record[1])
	}
	return websiteList
}
