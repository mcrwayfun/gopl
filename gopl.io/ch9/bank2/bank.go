package main

import "sync"

var (
	sema    = make(chan struct{}, 1)
	balance int
)

func Deposit(amount int) {
	sema <- struct{}{}
	balance = balance + amount
	<-sema
}

func Balance() int {
	sema <- struct{}{}
	b := balance
	<-sema
	return b
}

func main() {
	var n sync.WaitGroup
	for i := 1; i <= 100; i++ { // total should be 5050
		n.Add(1)
		go func(amount int) {
			Deposit(amount)
			n.Done()
		}(i)
	}
	n.Wait()

	balance := Balance()
	println(balance)
}
