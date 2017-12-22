package main

import (
	"bytes"
	"fmt"
)

func main() {
	a := []int{1, 4, 7, 9, 98}
	fmt.Println(intsToString(a))

}

func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()

}
