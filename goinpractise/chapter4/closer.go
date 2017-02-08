package main

import "fmt"

func main() {
	defer func() {
		//		msg := 999
		fmt.Println(msg)
	}()
	msg = 888
	fmt.Println(msg)
}
