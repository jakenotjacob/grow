package main

import (
	"fmt"
	"github.com/adonovan/gopl.io/ch5/links"
	"log"
	"os"
)

// Create counting semaphore to limit concurrent requests
var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)
	// Aquire token
	tokens <- struct{}{}
	list, err := links.Extract(url)
	// Release token
	<-tokens
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	worklist := make(chan []string)
	// number of pending sends to worklist
	var n int

	n++
	go func() { worklist <- os.Args[1:] }()

	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}
