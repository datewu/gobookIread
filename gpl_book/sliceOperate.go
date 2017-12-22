package main

func main() {

}

// nonEmpty returns a slice holding onlu the non-empty string.
// The underlying array is modified during the call.
func nonEmpty(sour []string) []string {
	var i int
	for _, s := range sour {
		if s != "" {
			sour[i] = s
			i++
		}
	}
	return sour[:i]
}

func nonEmptyIntact(sour []string) []string {
	out := sour[:0]
	for _, s := range sour {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

// reserve the order
func removeAt(a []int, i int) []int {
	copy(a[i:], a[i+1:])
	return a[:len(a)-1]
}

// does NOT reserve the order
func removeAt2(a []int, i int) []int {
	a[i] = a[len(a)-1]
	return a[:len(a)-1]
}
