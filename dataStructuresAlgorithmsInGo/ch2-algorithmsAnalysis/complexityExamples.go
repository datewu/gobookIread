package main

// o(n + n-1 + n-2 + ... + 1) = o(n2)
func fun3(count int) int {
	m := 0
	for i := 0; i < count; i++ {
		for j := 0; j < i; j++ {
			m++

		}

	}
	return m
}

// o(log(n))
func fun4(count int) int {
	m := 0
	for i := 0; i < count; i *= 2 {
		m++
	}
	return m
}

// o(n)
func fun12(count int) int {
	m := 0
	j := 0
	for i := 0; i < count; i++ {
		for ; j < count; j++ {
			m++
		}
	}
	return m
}

// o( 1 + 2 + 4 + 8 + 16 + n/2 +n) = o(n)
func fun13(count int) int {
	m := 0
	for i := 0; i < count; i *= 2 {
		for j := 0; j <= i; j++ {
			m++
		}
	}
	return m
}
