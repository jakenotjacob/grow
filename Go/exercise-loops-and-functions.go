package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := x / 2
	for y := 0; y <= 10; y++ {
		if z*z == x {
			fmt.Println("Iterations: ", y)
			return z
		} else if z == z - ((z*z - x) / (2 * z)) {
			fmt.Println("Repeat after iteration: ", y)
			return z
		}
		
		z -= (z*z - x) / (2 * z)

		fmt.Println(z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(9))
}

