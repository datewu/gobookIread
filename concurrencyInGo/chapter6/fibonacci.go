package main

import "fmt"

func main() {

	var fib func(n int) <-chan int

	fib = func(n int) <-chan int {
		result := make(chan int)
		go func() {
			defer close(result)
			if n <= 2 {
				result <- 1
				return
			}
			result <- <-fib(n-1) + <-fib(n-2)

		}()

		return result
	}

	fmt.Println("fib(4) =", <-fib(4))

}
