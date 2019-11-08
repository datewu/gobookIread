package main

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

func main() {
	myV := runtime.Version()
	a := strings.Split(myV, ".")
	major := a[0][2]
	minor := a[1]
	m1, _ := strconv.Atoi(string(major))
	m2, _ := strconv.Atoi(minor)

	if m1 == 1 && m2 < 8 {
		fmt.Println("need Go version or higher!")
		return
	}

	fmt.Println("you are using Go bersion", myV, major, minor)
}

// ➜  chapter2 git:(master) ✗ go run require_version.go
// you are using Go bersion go1.13.3 49 13
