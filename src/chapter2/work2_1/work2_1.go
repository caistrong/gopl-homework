package main

import (
	"fmt"

	"github.com/caistrong/gopl-homework/src/chapter2/work2_1/tempconv"
)

func main() {
	var k tempconv.Kelvin
	fmt.Printf("Kelvin: %s\tCelsius: %s\t Fahrenheit: %s\n", k.String(), tempconv.KToC(k).String(), tempconv.KToF(k).String())
}
