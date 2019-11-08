package main

import (
	"fmt"
	"runtime"
	"time"
)

func printStats(mem runtime.MemStats) {
	runtime.ReadMemStats(&mem)
	fmt.Println("mem.Alloc:", mem.Alloc)
	fmt.Println("mem.TotalAlloc:", mem.TotalAlloc)
	fmt.Println("mem.HeapAlloc:", mem.HeapAlloc)
	fmt.Println("mem.NumGC:", mem.NumGC)
	fmt.Println("-----")
}

func main() {
	var m runtime.MemStats
	printStats(m)
	for index := 0; index < 10; index++ {
		s := make([]byte, 500000000)
		if s == nil {
			fmt.Println("Operation failed!")
		}

	}
	printStats(m)
	for index := 0; index < 10; index++ {
		s := make([]byte, 5000000000)
		if s == nil {
			fmt.Println("Operation failed!")
		}
		time.Sleep(5 * time.Second)
	}
	printStats(m)
}
