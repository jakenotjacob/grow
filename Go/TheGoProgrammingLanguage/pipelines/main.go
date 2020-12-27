package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		defer close(naturals)
		for x := 0; x < 100; x++ {
			naturals <- x
		}
	}()

	// Squarer
	go func() {
		defer close(squares)
		for {
			x, ok := <-naturals
			if !ok {
				break // Channel was closed and drained
			}
			squares <- x * x
		}
	}()

	// Printer
	for x := range squares {
		fmt.Println(x)
	}
}
