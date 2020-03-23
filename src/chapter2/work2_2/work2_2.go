package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/caistrong/gopl-homework/src/chapter2/work2_2/lenconv"
)

// run command:
// go run ./src/chapter2/work2_2/work2_2.go 12
// OR
// go run ./src/chapter2/work2_2/work2_2.go
// 12
func main() {
	var ft lenconv.Foot
	if len(os.Args) == 1 {
		fmt.Println("please input a number (unit ft):")
		var input float64
		fmt.Scanf("%g", &input)
		ft = lenconv.Foot(input)
		fmt.Printf("%s = %s = %s\n", ft, lenconv.FToM(ft).String(), lenconv.FToI(ft).String())

	} else {
		for _, inputStr := range os.Args[1:] {
			f, err := strconv.ParseFloat(inputStr, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "err: %v", err)
			}

			ft = lenconv.Foot(f)

			fmt.Printf("%s = %s = %s\n", ft, lenconv.FToM(ft).String(), lenconv.FToI(ft).String())
		}
	}
}
