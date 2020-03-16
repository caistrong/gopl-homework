package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	startTime := time.Now()
	dumpPrint()
	endTime := time.Now()
	fmt.Println("dumpPrint time cost: " + endTime.Sub(startTime).String())

	startTime2 := time.Now()
	joinPrint()
	endTime2 := time.Now()
	fmt.Println("joinPrint time cost: " + endTime2.Sub(startTime2).String())

}

func dumpPrint() {
	s, sep := "", ""
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func joinPrint() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
