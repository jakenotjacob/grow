package main

import "C"

// Build package as C shared lib
// go build -buildmode=c-shared -o ohno.so

//Comment below to export the function symbol

//export Boop
func Boop(s string) string {
	return s
}

func main() {}
