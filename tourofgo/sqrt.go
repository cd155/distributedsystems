package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for i:=0; i<10;i++{
		head := (z*z - x) / (2*z)

		if math.Abs(head) < 0.0000001{
			fmt.Println(i)
			return z
		}
		z -= head
	}
	return z
}

func main() {
	fmt.Println("My Sqrt:", Sqrt(2))
	fmt.Println("Go Sqrt:", math.Sqrt(2))
}
