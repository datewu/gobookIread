package main

import (
	"runtime"
)

func main() {
	var n = 36_000_000
	m := make([]map[int]int, 200)
	for i := range m {
		m[i] = make(map[int]int)
	}
	for index := 0; index < n; index++ {
		v := int(index)
		m[index%200][v] = v

	}
	runtime.GC()
	_ = m[0][0]
	//	fmt.Println(s[0])
}

// ➜  chapter2 git:(master) ✗ time go run map_split.go
// go run map_split.go  11.26s user 0.73s system 100% cpu 11.910 total
