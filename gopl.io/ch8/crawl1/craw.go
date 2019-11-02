package main

import (
	"../links"
	"fmt"
	"log"
	"os"
)

func main() {
	worklist := make(chan []string)

	go func() {// 接收命令行输入的参数
		worklist <- os.Args[1:]
	}()

	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				go func(link string) {
					worklist <- crawl(link)
				}(link)// 显示传入,避免循环变量快照
			}
		}
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
