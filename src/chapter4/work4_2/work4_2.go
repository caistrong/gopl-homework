package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

// HashAlgo Const
const (
	SHA256 = "SHA256"
	SHA384 = "SHA384"
	SHA512 = "SHA512"
)

// run command
// go run ./src/chapter4/work4_2/work4_2.go caixiaocong
// go run ./src/chapter4/work4_2/work4_2.go -algo SHA384 caixiaocong
// go run ./src/chapter4/work4_2/work4_2.go -algo SHA512 caixiaocong
func main() {
	var algo string
	flag.StringVar(&algo, "algo", "SHA256", "USAGE: available value is SHA256（default）or SHA384 or SHA512")
	flag.Parse()
	for _, s := range flag.Args() {
		resultStr := ""
		switch algo {
		case SHA256:
			c := sha256.Sum256([]byte(s))
			resultStr = fmt.Sprintf("%x", c)
		case SHA384:
			c := sha512.Sum384([]byte(s))
			resultStr = fmt.Sprintf("%x", c)
		case SHA512:
			c := sha512.Sum512([]byte(s))
			resultStr = fmt.Sprintf("%x", c)
		default:
			fmt.Printf("Error Hash Algo Type %s!! available value is SHA256（default）or SHA384 or SHA512\n", algo)
			os.Exit(1)
		}
		fmt.Printf("str:%s\talgo:%s\t sha: %s\n", s, algo, resultStr)
	}
}
