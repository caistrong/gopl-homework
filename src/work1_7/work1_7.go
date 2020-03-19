package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// run command:
// go run work1_7.go http://info.cern.ch/hypertext/WWW/TheProject.html
func main() {
	for _, url := range os.Args[1:] {
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
