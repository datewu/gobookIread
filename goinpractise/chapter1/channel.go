package main

import (
	"fmt"
	"time"
)

func printCount(c chan int) {
	num := 0
	for num >= 0 {
		num = <-c
		fmt.Print(num, " ")
	}
}

func main() {
	c := make(chan int)
	a := []int{2, 5, 9, 3, 0, 4, -4}
	go printCount(c)

	for _, v := range a {
		c <- v
	}
	time.Sleep(399 * time.Millisecond)
	fmt.Println("End of main")
	fmt.Println(a)

}
