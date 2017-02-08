package main

import (
	"fmt"
	"log"
)

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance
var withdrwas = make(chan ws) // send amount to deposit

type ws struct {
	amount   int
	callback chan bool
}

func Deposit(amount int) {
	deposits <- amount
}

func Balance() int {
	return <-balances
}

func Withdraw(amount int) (r bool) {
	w := ws{amount, make(chan bool)}
	withdrwas <- w
	r = <-w.callback
	log.Println(r)
	return
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
			log.Println("sb is checking his balance")
		case w := <-withdrwas:
			if w.amount <= balance {
				balance -= w.amount
				w.callback <- true
				continue
			}
			w.callback <- false
		}
	}
}

func init() {
	go teller()

}
func main() {
	fmt.Println(Balance())
	Deposit(100)
	fmt.Println(Balance())
	Withdraw(20)
	fmt.Println(Balance())
}
