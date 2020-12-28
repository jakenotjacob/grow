package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	files := os.Args[1:]
	fmt.Println(strings.Join(files, "-"))
}
