package main

import "fmt"

// Storage Unit
const (
	KB = 1000
	MB = KB * KB
	GB = MB * KB
	TB = GB * KB
	PB = TB * KB
	EB = PB * KB
	ZB = EB * KB
	YB = ZB * KB
)

// run command
// go run ./src/chapter3/work3_13/work3_13.go
func main() {
	// checkt default Type for const value
	// 画蛇添足地查看无类型常量的默认类型
	fmt.Printf("%T\n", 0)
	fmt.Printf("%T\n", 0.0)
	fmt.Printf("%T\n", 0i)
	fmt.Printf("%T\n", '\000') // "int32" (rune)
}
