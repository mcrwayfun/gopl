package main

import (
	"../../../gopl.io/ch8/links"
	"fmt"
	"log"
	"os"
	"sync"
)

var tokens = make(chan struct{}, 20)
var maxDepth int
var seen = make(map[string]bool)
var seenLock = sync.Mutex{}

func crawl(url string, depth int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(depth, url)
	if depth >= maxDepth {
		return
	}
	tokens <- struct{}{} // start
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	for _, link := range list {
		seenLock.Lock() // 加锁,防止并发访问同个key
		if seen[link] {
			seenLock.Unlock()
			continue
		}
		seen[link] = true
		seenLock.Unlock()
		wg.Add(1)
		go crawl(link, depth+1, wg)
	}
}

func main() {
	maxDepth = 3
	wg := &sync.WaitGroup{}
	for _, link := range os.Args[1:] {
		wg.Add(1)
		go crawl(link, 0, wg)
	}
	wg.Wait()
}
