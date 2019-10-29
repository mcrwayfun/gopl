package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	//defer conn.Close()
	done := make(chan struct{}) // channel使用make初始化，用struct来构造表示只用来传递消息
	go func() { // 模拟异步操作
		_, _ = io.Copy(os.Stdout, conn)
		log.Println("done")
		done <- struct{}{} // 发送消息给main goroutine
	}()
	mustCopy(conn, os.Stdin)
	conn.Close()
	<- done
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
