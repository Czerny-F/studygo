package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := x / 2
	for i := 0; i < 10; i++ {
		t := z - (z*z-x)/(2*z)
		fmt.Printf("Processing: %g\n", t)
		if t == z {
			return z
		}
		z = t
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
