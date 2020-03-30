package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/caistrong/gopl-homework/src/chapter4/work4_12/xkcd"
)

//NOTE: 原题要求使用 https://omdbapi.com/ 下载电影海报，不过该API目前要付费，最低配1$/Month，没钱买啊所以这里沿用上一题的API

// run command
// go run ./src/chapter4/work4_13/work4_13.go Privacy Tradition
func main() {
	// xkcd.GenerateFile()//需要生成xkcd.json文件时去掉这行注释
	for _, queryWord := range os.Args[1:] {
		workDir, _ := os.Getwd()
		outputFilePath := workDir + "/src/chapter4/work4_12/xkcd.json"
		xkcdFile, err := os.Open(outputFilePath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "opening error:%v", err)
		}

		comicMetaSlice := []xkcd.ComicMeta{}
		if err = json.NewDecoder(xkcdFile).Decode(&comicMetaSlice); err != nil {
			fmt.Fprintf(os.Stderr, "decode error:%v", err)
		}
		comicMeta := findByTitle(comicMetaSlice, queryWord)
		saveImageFromURL(comicMeta.Img, workDir+"/src/chapter4/work4_13/images/")
		fmt.Printf("%s: %s\n", queryWord, comicMeta.Img)
	}
}

func findByTitle(s []xkcd.ComicMeta, title string) xkcd.ComicMeta {
	for _, i := range s {
		if title == i.Title {
			return i
		}
	}
	return xkcd.ComicMeta{}
}

func saveImageFromURL(url, distDir string) error {
	// don't worry about errors
	response, e := http.Get(url)
	if e != nil {
		return e
	}
	defer response.Body.Close()

	//open a file for writing
	file, err := os.Create(distDir + filepath.Base(url))
	if err != nil {
		return e
	}
	defer file.Close()

	// Use io.Copy to just dump the response body to the file. This supports huge files
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}
	return nil
}
