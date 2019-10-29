package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

type clock struct {
	name, host string
}

func main() {
	if len(os.Args) == 1 { // 没有输入任何参数
		_, _ = fmt.Fprintln(os.Stderr, "usage: NAME:HOST")
		os.Exit(1)
	}
	clocks := make([]*clock, 0)
	for _, v := range os.Args[1:] {
		fields := strings.Split(v, "=")
		if len(fields) != 2 {
			_, _ = fmt.Fprintf(os.Stderr, "bad args: %s\n", v)
			os.Exit(1)
		}
		clocks = append(clocks, &clock{fields[0], fields[1]})
	}

	for _, c := range clocks {
		conn, err := net.Dial("tcp", c.host)
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()
		go c.watch(os.Stderr, conn)
	}
	for {
		time.Sleep(time.Minute)
	}
}

func (c *clock) watch(w io.Writer, r io.Reader) {
	s := bufio.NewScanner(r)
	for s.Scan() {
		_, _ = fmt.Fprintf(w, "%s: %s\n", c.name, s.Text())
	}
	fmt.Println(c.name, "done")
	if s.Err() != nil {
		log.Printf("can't read from %s: %s", c.name, s.Err())
	}
}
