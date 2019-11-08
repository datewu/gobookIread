package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		panic("not enough arguments!")
	}
	fmt.Println("thanks for the argument(s)!")
}
