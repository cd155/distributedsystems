package main

import "fmt"

var a int = 20
const b = 1 << 3

func main() {
	const Pi = 3.14
	okie := "ok"
	fmt.Println("Hello, 世界", a, okie, Pi, b)
}
