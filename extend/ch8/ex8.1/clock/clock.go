package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

// 使用 -port 指定端口
var port = flag.Int("port", 8080, "listen port")

func main() {
	flag.Parse()
	log.Print(*port)
	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatal(err)
	}

	// 死循环
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}
