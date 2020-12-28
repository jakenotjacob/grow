package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	count := make(map[string]int)

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		count[input.Text()]++
	}
	fmt.Println(count)
}
