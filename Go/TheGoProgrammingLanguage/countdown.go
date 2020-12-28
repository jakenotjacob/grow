package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	abort := make(chan struct{})
	go func() {
		// read single byte
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	fmt.Println("BEGIN THE COUNT!")
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		select {
		case <-tick:
			fmt.Println(countdown)
		case <-abort:
			fmt.Println("Launch aborted!")
			return
		}
	}
}

// launch()
