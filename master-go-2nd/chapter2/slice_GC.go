package main

import (
	"runtime"
)

type d struct {
	i, j int
}

func main() {
	var n = 36_000_000
	var s []d
	for index := 0; index < n; index++ {
		v := int(index)
		s = append(s, d{v, v})
	}
	runtime.GC()
	_ = s[0]
	//	fmt.Println(s[0])
}

// ➜  chapter2 git:(master) ✗ time go run slice_GC.go
// go run slice_GC.go  1.40s user 0.61s system 157% cpu 1.277 total
