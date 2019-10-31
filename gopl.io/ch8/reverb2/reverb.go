package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
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

func echo(c net.Conn, shout string, delay time.Duration, wg sync.WaitGroup) {
	_, _ = fmt.Fprintln(c, "\t", strings.ToUpper(shout)) // 接收到的输入输出到conn
	time.Sleep(delay)
	_, _ = fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	_, _ = fmt.Fprintln(c, "\t", strings.ToLower(shout))
	wg.Done() // 计数完成,结束
}

func handleConn(c net.Conn) {
	var wg sync.WaitGroup
	scanner := bufio.NewScanner(c)
	wg.Add(1)
	for scanner.Scan() { // 存在输入
		echo(c, scanner.Text(), 1*time.Second, wg)
	}
	wg.Wait() // 等待结束才关闭连接
	c.Close()
}
