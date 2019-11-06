package main

import "fmt"

type Withdrawal struct {
	amount  int
	success chan bool
}

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance
var withdrawals = make(chan Withdrawal)

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }
func Withdraw(amount int) bool {
	ch := make(chan bool)
	withdrawals <- Withdrawal{amount, ch}
	return <-ch
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case w := <-withdrawals:
			if w.amount > balance {
				w.success <- false
				continue
			}
			balance -= w.amount
			w.success <- true
		case balances <- balance:
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}

func main(){
	b1 := Balance()
	ok := Withdraw(50)
	if !ok {
		fmt.Printf("ok = false, want true. balance = %d\n", Balance())
	}
	expected := b1 - 50
	if b2 := Balance(); b2 != expected {
		fmt.Printf("balance = %d, want %d\n", b2, expected)
	}
}