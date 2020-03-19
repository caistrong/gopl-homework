package main

import (
	"bufio"
	"fmt"
	"os"
)

type countKey struct {
	line     string
	fileName string
}

// run command:
// go run work1_4.go test.txt
func main() {
	counts := make(map[countKey]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, file := range files {
			f, err := os.Open(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "open file %s fail, err : %v\n", file, err)
			}
			countLines(f, counts)
			f.Close()
		}
	}
	fmt.Println("出现次数\t行内容\t文件名")
	for cntKey, cnt := range counts {
		if cnt > 1 {
			fmt.Printf("%d\t%s\t%s\n", cnt, cntKey.line, cntKey.fileName)
		}
	}

}

func countLines(f *os.File, counts map[countKey]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		countKey := countKey{
			line:     input.Text(),
			fileName: f.Name(),
		}
		counts[countKey]++
	}
}
