package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
)

func main() {
	flag.Parse()
	cancel := make(chan struct{})
	responses := make(chan *http.Response)
	wg := &sync.WaitGroup{}
	for _, url := range flag.Args() {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			req, err := http.NewRequest("GET", url, nil)
			if err != nil {
				log.Printf("GET %s: %s", url, err)
				return
			}
			req.Cancel = cancel
			response, err := http.DefaultClient.Do(req)
			if err != nil {
				log.Printf("GET %s: %s", url, err)
				return
			}
			responses <- response
		}(url)
	}
	resp := <-responses
	defer resp.Body.Close()
	close(cancel)// close when the first request finish
	fmt.Println(resp.Request.URL)
	for name, vals := range resp.Header {
		fmt.Printf("%s: %s\n", name, strings.Join(vals, ","))
	}

	wg.Wait()
}
