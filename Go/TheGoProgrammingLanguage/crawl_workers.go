package main

import (
	"os"
)

func main() {
	worklist := make(chan []string)
	unseenLinks := make(chan string)

	go func() { worklist <- os.Args[1:] }()

	// Create 20 crawler goroutines to fetch unseen links
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link)
				go func() { worklist <- foundLinks }()
			}
		}()
	}

	// Dedup worklist items and send unseen to crawlers
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
			}
		}
	}
}
