package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pengji/go-code/programming_general/func/links"
)

// crawl3 crawls web links starting with the command-line arguments.
//
// This version uses bounded parallelism.
// For simplicity, it does not address the termination problem.

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	worklist := make(chan []string)  //lists of urls, may have duplicates
	unseenLinks := make(chan string)  //de-duplicated urls

	// add command-line arguments to worklist.
	go func() {worklist <- os.Args[1:]} ()

	// create 20 crawler goroutines to fetch each unseen link.
	for i := 0; i<20; i++ {
		go func(){
			for link := range unseenLinks {
				foundLinks := crawl(link)
				go func() {worklist <- foundLinks}()
			}
		}()
	}

	// the main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers.
	seen := make(map[string] bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
			}
		}
	}
}