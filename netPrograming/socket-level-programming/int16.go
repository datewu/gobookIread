package main

import "fmt"

func main() {
	var n uint16 = 65534
	x := byte(n >> 8)
	y := byte(n & 255)
	z := rune(n) & 0xffff
	fmt.Println(x, y, z, n)
}
