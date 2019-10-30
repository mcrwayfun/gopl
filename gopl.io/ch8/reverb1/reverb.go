package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func echo(c net.Conn, shout string, delay time.Duration) {
	_, _ = fmt.Fprintln(c, "\t", strings.ToUpper(shout)) // 接收到的输入输出到conn
	time.Sleep(delay)
	_, _ = fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	_, _ = fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	scanner := bufio.NewScanner(c)
	for scanner.Scan() { // 存在输入
		echo(c, scanner.Text(), 1*time.Second)
	}
	c.Close()
}
