package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

// run command:
// go run ./src/chapter1/work1_10/work1_10.go http://qq.com http://qq.com http://qq.com http://qq.com http://qq.com http://qq.com
func main() {
	startTime := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}

	for range os.Args[1:] {
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
