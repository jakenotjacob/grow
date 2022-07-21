package hello

import (
	"fmt"

	"rsc.io/quote/v3"
)

func Hello() string {
	fmt.Println("asdf")
	return quote.HelloV3()
}

func Proverb() string {
	return quote.Concurrency()
}
