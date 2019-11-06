package main

import "fmt"

var deposits = make(chan int)
var balances = make(chan int)

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }

func teller() {
	var balance int
	for {
		select {
		case amount := <-deposits:
			balance += amount
			fmt.Printf("amount:%d balance:%d\n", amount, balance)
		case balances <- balance:
			fmt.Printf("%d balances <- balance", balance)
		}
	}
}

func init() {
	go teller()
}

func main(){
	done := make(chan struct{})

	// Alice
	go func() {
		Deposit(200)
		fmt.Println("=", Balance())
		done <- struct{}{}
	}()

	// Bob
	go func() {
		Deposit(100)
		done <- struct{}{}
	}()

	// Wait for both transactions.
	<-done
	<-done
}