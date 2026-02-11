package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string{
	return "cannot Sqrt negative number"
}

func Sqrt(x float64) (float64, error) {
	z := 1.0
	
	if x < 0{
		return x, ErrNegativeSqrt(x)
	}
	
	for i:=0; i<10;i++{
		head := (z*z - x) / (2*z)

		if math.Abs(head) < 0.0000001{
			fmt.Println(i)
			return z, nil
		}
		z -= head
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(-2))
}
