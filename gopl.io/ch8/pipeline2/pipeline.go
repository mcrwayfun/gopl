package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		for x := 0; ; x++ {
			naturals <- x
		}
	}()

	// Squarer
	go func() {
		for x := range naturals { // 使用range来判断channel中是否还有数据，若没有则直接跳出循环
			squares <- x * x
		}
		close(naturals)
	}()

	// Printer
	for x := range squares {
		fmt.Println(x)
	}
}
