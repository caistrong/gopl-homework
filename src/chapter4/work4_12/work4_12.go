package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/caistrong/gopl-homework/src/chapter4/work4_12/xkcd"
)

// run command
// go run ./src/chapter4/work4_12/work4_12.go Privacy Tradition
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
