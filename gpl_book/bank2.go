package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

var (
	sema    = make(chan struct{}, 1) // a binary semaphore guarding balance
	balance int
)

func Deposit(amonut int) {
	sema <- struct{}{} // acquire token
	balance += amonut
	<-sema // relaese token
}

func Balance() int {
	sema <- struct{}{} // acquire token
	b := balance
	<-sema // relaese token
	return b
}
