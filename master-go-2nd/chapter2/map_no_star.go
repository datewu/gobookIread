package main

import (
	"runtime"
)

func main() {
	var n = 36_000_000
	m := make(map[int]int)
	for index := 0; index < n; index++ {
		v := int(index)
		m[v] = v
	}
	runtime.GC()
	_ = m[0]
	//	fmt.Println(s[0])
}

// ➜  chapter2 git:(master) ✗ time go run map_no_star.go
// go run map_no_star.go  11.61s user 1.14s system 100% cpu 12.707 total
