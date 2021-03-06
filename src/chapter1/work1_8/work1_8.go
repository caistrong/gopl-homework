package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

// run command:
// go run ./src/chapter1/work1_8/work1_8.go http://info.cern.ch/hypertext/WWW/TheProject.html
func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "get %s fail, err: %v", url, err)
			os.Exit(1)
		}
		_, copyErr := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if copyErr != nil {
			fmt.Fprintf(os.Stderr, "io.Copy err: %v", copyErr)
			os.Exit(1)
		}
	}
}
