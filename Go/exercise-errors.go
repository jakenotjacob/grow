package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string{
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return x, ErrNegativeSqrt(x)
	}
	
	z := x / 2
	for y := 0; y <= 10; y++ {
		if z*z == x {
			fmt.Println("Iterations: ", y)
			return z, nil
		} else if z == z - ((z*z - x) / (2 * z)) {
			fmt.Println("Repeat after iteration: ", y)
			return z, nil
		}
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}

