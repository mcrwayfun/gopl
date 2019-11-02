package main

import (
	"../links"
	"fmt"
	"log"
	"os"
)

func main() {
	worklist := make(chan []string)
	unseenlinks := make(chan string)

	go func() {
		worklist <- os.Args[1:]
	}()

	for i := 0; i < 20; i++ {
		go func() {
			fmt.Println("foundLinks pre")
			for link := range unseenlinks {
				foundLinks := crawl(link)
				go func() { worklist <- foundLinks }()
			}
		}()
	}

	seen := make(map[string]bool)
	for list := range worklist {
		fmt.Println("worklist pre")
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenlinks <- link
			}
		}
		fmt.Println("worklist finish")
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}
