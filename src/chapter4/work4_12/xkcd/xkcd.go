package xkcd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

// XkcdHost is host of https://xkcd.com/571/info.0.json
const XkcdHost = "https://xkcd.com/"

// ComicMeta is meta of comic
type ComicMeta struct {
	Month      string
	Num        int
	Link       string
	Year       string
	News       string
	SafeTitle  string `json:"safe_title"`
	Transcript string
	Img        string
	Title      string
	Day        string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

const (
	maxGoroutines = 5 //最多同时有maxGoroutines个http请求在处理
	comicCounts   = 2287
)

// GenerateFile generate xkcd.json store the comic data
func GenerateFile() {
	comicMetaSlice := []ComicMeta{}
	cmch := make(chan ComicMeta)
	guard := make(chan struct{}, maxGoroutines)
	startTime := time.Now()
	go makeRequest(guard, cmch)
	for i := 1; i < comicCounts; i++ {
		comicMetaSlice = append(comicMetaSlice, <-cmch)
	}
	costTime := time.Since(startTime).Seconds()
	fmt.Printf("cost %.2f s\n", costTime)
	text, err := json.Marshal(comicMetaSlice)
	check(err)
	workDir, wdErr := os.Getwd()
	check(wdErr)
	outputFilePath := workDir + "/src/chapter4/work4_12/xkcd.json"
	fmt.Println("outputFilePath:", outputFilePath)
	writeErr := ioutil.WriteFile(outputFilePath, text, 0644)
	check(writeErr)
}

func makeRequest(guard chan struct{}, cmch chan ComicMeta) {
	for i := 1; i < comicCounts; i++ {
		guard <- struct{}{} // would block if guard channel is already filled
		go func(n int) {
			getComicMetaByID(n, cmch)
			<-guard
		}(i)
	}
}

func getComicMetaByID(cid int, cmch chan ComicMeta) error {
	finalURL := XkcdHost + strconv.Itoa(cid) + "/info.0.json"
	fmt.Println("getting...", finalURL)
	resp, err := http.Get(finalURL)
	fmt.Println("resp.Status...", resp.Status)
	check(err)
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		cmch <- ComicMeta{}
		return fmt.Errorf("http response errot, resp.Status == %s", resp.Status)
	}
	var comicMeta ComicMeta
	if err := json.NewDecoder(resp.Body).Decode(&comicMeta); err != nil {
		resp.Body.Close()
		cmch <- ComicMeta{}
		return err
	}
	resp.Body.Close()
	cmch <- comicMeta // would block if guard channel is already filled
	return nil
}
