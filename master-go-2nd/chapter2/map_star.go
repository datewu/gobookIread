package main

import (
	"runtime"
)

func main() {
	var n = 36_000_000
	m := make(map[int]*int)
	for index := 0; index < n; index++ {
		v := int(index)
		m[v] = &v
	}
	runtime.GC()
	_ = m[0]
	//	fmt.Println(s[0])
}

// ➜  chapter2 git:(master) ✗ time go run map_star.go
// go run map_star.go  25.89s user 1.53s system 210% cpu 13.020 total
